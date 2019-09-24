import xml.etree.ElementTree as ET
import urllib2
import pandas as pd
import pythondbtools.dbtools

"""
id              (UUID)
full_name       (None) (wholeName)
type            (None) (regulationType)
source          (None) (publicationUrl)
address         (None) (city, zipcode)
updated_at      (Now)  (publicationDate)
created_at      (Now)  ()
"""


def parse_file():
    # Variables that contain the credentials and URLS to acccess the data
    DOWNLOAD_URL = "https://webgate.ec.europa.eu/europeaid/fsd/fsf/public/files/xmlFullSanctionsList_1_1/content?token=n002wvni"
    file = urllib2.urlopen(DOWNLOAD_URL)
    data = file.read()
    file.close()
    tree = ET.fromstring(data)
    return tree


def select_fields(tree):
    wholeName = []
    type = []
    source = []
    address = []
    for sanctionEntity in tree.findall(
        "{http://eu.europa.ec/fpi/fsd/export}sanctionEntity"
    ):
        nameAlias = sanctionEntity.find("{http://eu.europa.ec/fpi/fsd/export}nameAlias")
        regulationSummary = nameAlias.find(
            "{http://eu.europa.ec/fpi/fsd/export}regulationSummary"
        )
        birthdate = sanctionEntity.find("{http://eu.europa.ec/fpi/fsd/export}birthdate")

        wholeName.append(nameAlias.attrib["wholeName"])
        type.append("SANCTION")
        source.append(regulationSummary.attrib["publicationUrl"])
        if birthdate is not None:
            address.append(
                birthdate.attrib["place"]
                + ", "
                + birthdate.attrib["city"]
                + ", "
                + birthdate.attrib["region"]
                + ", "
                + birthdate.attrib["zipCode"]
            )
        else:
            address.append("None")

        # returned_cols = ["full_name", "type", "source", "address"]

    sanctions_df = pd.DataFrame(
        {"full_name": wholeName, "type": type, "source": source, "address": address}
    )

    return sanctions_df


def remove_sanction_from_db():
    queried_bad_persons_list = pythondbtools.dbtools.read_bad_persons()
    for bad_person in queried_bad_persons_list:
        if bad_person.type == "SANCTION":
            pythondbtools.dbtools.delete_bad_person(bad_person)


def add_new_sanction_to_db(df):
    pythondbtools.dbtools.push_df(df)


def run(event, context):
    tree = parse_file()
    data = select_fields(tree)
    remove_sanction_from_db()
    add_new_sanction_to_db(data)
    return {
        "statusCode": 200,
        "body": "<html><body><p>Added SANCTION list to database.</p></body></html>",
        "headers": {"Content-Type": "text/html"},
    }


if __name__ == "__main__":
    tree = parse_file()
    data = select_fields(tree)
    remove_sanction_from_db()
    add_new_sanction_to_db(data)
