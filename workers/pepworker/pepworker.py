import re
import traceback
from json import dumps

import requests
from bs4 import BeautifulSoup

from pandas import read_excel
from workers.pythondbtools import dbtools
from workers.pythondbtools.dbtools import BAD_PERSON_TYPE


def get_link_to_file():
    base_url = "https://www.finanstilsynet.dk"
    page_url = "/Tal-og-Fakta/PEP-liste"
    req = requests.get(f"{base_url}{page_url}")
    soup = BeautifulSoup(req.content, "html.parser")
    file_href = soup.find_all(href=re.compile("PEPliste"))
    file_url = file_href[0].get("href")  # do test here for multiple links
    return f"{base_url}{file_url}"


def parse_pep_xlsx(link):
    data = read_excel(link, header=1, index_col=None).reset_index(drop=True)
    parsing_cols = [
        "Navn",
        "Unnamed: 2",
        "Stillingsbetegnelse ",
        "Tilføjet på PEP-listen (dato)",
    ]
    rename_dict = {
        "Navn": "sur_name",
        "Unnamed: 2": "first_name",
        "Stillingsbetegnelse ": "occupation",
        "Tilføjet på PEP-listen (dato)": "date_added",
    }
    returned_cols = ["full_name", "type", "source", "country_code"]

    names_df = data[parsing_cols]
    names_df = names_df.dropna(axis="index")
    names_df = names_df.rename(columns=rename_dict)
    names_df["type"] = BAD_PERSON_TYPE.PEP
    names_df["source"] = link
    names_df["country_code"] = "DK"
    names_df["full_name"] = names_df["first_name"] + " " + names_df["sur_name"]
    names_df = names_df[returned_cols]
    names_df.drop_duplicates(keep="first", inplace=True)
    names_df["country_code"] = names_df["country_code"].apply(lambda x: [x])
    return names_df


def remove_pep_from_db():
    dbtools.delete_all_bad_persons(list_type=BAD_PERSON_TYPE.PEP)


def add_new_pep_to_db(df):
    dbtools.update_df(df, list_type=BAD_PERSON_TYPE.PEP)


def upsert_new_pep_to_db(df):
    result_string = dbtools.upsert_df(df, list_type=BAD_PERSON_TYPE.PEP)
    return result_string


def run(event, context):
    try:
        file_link = get_link_to_file()
        parsed_df = parse_pep_xlsx(file_link)
        add_new_pep_to_db(parsed_df)
        return {
            "statusCode": 200,
            "body": '{"data": "pepworker finished"}',
            "headers": {"Content-Type": "application/json"},
        }
    except Exception as err:
        body_dict = {
            "error": "pepworker failed",
            "error message": traceback.format_exc().split("\n"),
        }
        return {
            "statusCode": 500,
            "body": dumps(body_dict),
            "headers": {"Content-Type": "application/json"},
        }


def upsert_run(event, context):
    try:
        file_link = get_link_to_file()
        parsed_df = parse_pep_xlsx(file_link)
        result_string = upsert_new_pep_to_db(parsed_df)
        body_dict = {
            "data": "pepworker finished",
            "message": result_string.split("\n"),
        }
        return {
            "statusCode": 200,
            "body": dumps(body_dict),
            "headers": {"Content-Type": "application/json"},
        }
    except Exception as err:
        body_dict = {
            "error": "pepworker failed",
            "error message": traceback.format_exc().split("\n"),
        }
        return {
            "statusCode": 500,
            "body": dumps(body_dict),
            "headers": {"Content-Type": "application/json"},
        }


if __name__ == "__main__":
    file_link = get_link_to_file()
    parsed_df = parse_pep_xlsx(file_link)
    upsert_new_pep_to_db(parsed_df)
