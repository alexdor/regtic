import xml.etree.ElementTree as ET
import urllib.request
import pandas as pd
import traceback
from workers.pythondbtools import dbtools
from workers.pythondbtools.dbtools import BAD_PERSON_TYPE, ADDRESS_TYPE
import os


def parse_file():
    try:
        access_token = os.environ["SANCTIONS_LIST"]
    except:
        access_token = "https://webgate.ec.europa.eu/europeaid/fsd/fsf/public/files/xmlFullSanctionsList_1_1/content?token=n002wvni"

    file = urllib.request.urlopen(access_token)
    data = file.read()
    file.close()
    return ET.fromstring(data)


def select_fields(tree):
    wholeName = []
    alias = []
    type = []
    source = []
    address_type = []
    citizenship_region = []
    citizenship_code = []
    street = []
    poBox = []
    city = []
    zipCode = []
    region = []
    place = []
    country = []

    subject_type = []
    sanctions_export = "{http://eu.europa.ec/fpi/fsd/export}"

    for sanctionEntity in tree.findall(sanctions_export + "sanctionEntity"):

        person_alias = []

        type_l = []
        street_l = []
        poBox_l = []
        city_l = []
        zipCode_l = []
        region_l = []
        place_l = []
        country_l = []

        country_arr = []
        citizenship_region_l = []

        nameAlias = sanctionEntity.findall(sanctions_export + "nameAlias")
        regulationSummary = nameAlias[0].find(sanctions_export + "regulationSummary")

        subjectType = sanctionEntity.find(sanctions_export + "subjectType")

        birthdate_list = sanctionEntity.findall(sanctions_export + "birthdate")

        citizenship_list = sanctionEntity.findall(sanctions_export + "citizenship")

        address_list = sanctionEntity.findall(sanctions_export + "address")

        wholeName.append(nameAlias[0].attrib["wholeName"])

        for i in birthdate_list:
            type_l.append(ADDRESS_TYPE.birthplace)
            street_l.append("")
            poBox_l.append("")
            city_l.append(i.attrib["city"])
            zipCode_l.append(i.attrib["zipCode"])
            region_l.append(i.attrib["region"])
            place_l.append(i.attrib["place"])
            if i.attrib["countryIso2Code"] == "00":
                country_l.append("ZZ")
            else:
                country_l.append(i.attrib["countryIso2Code"])

        for i in address_list:
            type_l.append(ADDRESS_TYPE.address)
            street_l.append(i.attrib["street"])
            poBox_l.append(i.attrib["poBox"])
            city_l.append(i.attrib["city"])
            zipCode_l.append(i.attrib["zipCode"])
            region_l.append(i.attrib["region"])
            place_l.append(i.attrib["place"])
            if i.attrib["countryIso2Code"] == "00":
                country_l.append("ZZ")
            else:
                country_l.append(i.attrib["countryIso2Code"])

        for i in citizenship_list:
            if i.attrib["countryIso2Code"] == "00" or i.attrib["countryIso2Code"] == "":
                country_arr.append("ZZ")
            else:
                country_arr.append(i.attrib["countryIso2Code"])

            citizenship_region_l.append(i.attrib["region"])

        for i in nameAlias:
            person_alias.append(i.attrib["wholeName"])

        alias.append(person_alias)
        address_type.append(type_l)
        citizenship_code.append(country_arr)
        citizenship_region.append(citizenship_region_l)
        street.append(street_l)
        poBox.append(poBox_l)
        city.append(city_l)
        zipCode.append(zipCode_l)
        region.append(region_l)
        place.append(place_l)
        country.append(country_l)

        type.append(BAD_PERSON_TYPE.SANCTION)
        subject_type.append(subjectType.attrib["classificationCode"])
        source.append(regulationSummary.attrib["publicationUrl"])

    sanctions_df = pd.DataFrame(
        {
            "address_type": address_type,
            "citizenship_code": citizenship_code,
            "citizenship_region": citizenship_region,
            "street": street,
            "poBox": poBox,
            "city": city,
            "zipCode": zipCode,
            "region": region,
            "place": place,
            "entity": subject_type,
            "country": country,
            "type": type,
            "full_name": wholeName,
            "alias": alias,
            "source": source,
        }
    )

    return sanctions_df


def remove_sanction_from_db():
    dbtools.delete_all_bad_persons(list_type=BAD_PERSON_TYPE.SANCTION)


def add_new_sanction_to_db(df):
    dbtools.update_df(df, list_type=BAD_PERSON_TYPE.SANCTION)


def run(event, context):
    try:
        tree = parse_file()
        data = select_fields(tree)
        add_new_sanction_to_db(data)
        return {
            "statusCode": 200,
            "body": '{"data": "sanctionworker finished"}',
            "headers": {"Content-Type": "application/json"},
        }
    except Exception as err:
        return {
            "statusCode": 500,
            "body": '{"error": "sanctionworker failed, error: '
            + traceback.format_exc()
            + '"}',
            "headers": {"Content-Type": "application/json"},
        }


if __name__ == "__main__":
    tree = parse_file()
    data = select_fields(tree)
    add_new_sanction_to_db(data)
