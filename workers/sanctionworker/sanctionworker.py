import os
import traceback
import urllib.request
import xml.etree.ElementTree as ET

import pandas as pd

from workers.pythondbtools import dbtools
from workers.pythondbtools.dbtools import BAD_PERSON_TYPE


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
    address = []
    sanctions_export = "{http://eu.europa.ec/fpi/fsd/export}"

    for sanctionEntity in tree.findall(sanctions_export + 'sanctionEntity'):
        nameAlias = sanctionEntity.findall(sanctions_export + 'nameAlias')
        # citizenship
        regulationSummary = nameAlias[0].find(sanctions_export + 'regulationSummary') #1
        birthdate = sanctionEntity.find(sanctions_export + 'birthdate') #1
        wholeName.append(nameAlias[0].attrib['wholeName'])
        person_alias = []
        for i in nameAlias:
            person_alias.append(i.attrib['wholeName'])

        alias.append(person_alias)
        type.append(BAD_PERSON_TYPE.SANCTION)
        source.append(regulationSummary.attrib['publicationUrl'])
        if birthdate is not None:
            address.append(
                f"{birthdate.attrib['place']}, {birthdate.attrib['city']}, {birthdate.attrib['region']}, {birthdate.attrib['zipCode']}"
            )
        else:
            address.append("None")

    sanctions_df = pd.DataFrame(
        {'full_name': wholeName,
         'alias': alias,
         'type': type,
         'source': source,
         'address': address
         })

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
