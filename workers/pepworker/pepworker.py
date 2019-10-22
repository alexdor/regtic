import os
import re

import pandas as pd
import requests
from bs4 import BeautifulSoup

from workers.pythondbtools import dbtools
from workers.pythondbtools.dbtools import TYPE_ENUM


def get_link_to_file():
    base_url = "https://www.finanstilsynet.dk"
    page_url = "/Tal-og-Fakta/PEP-liste"
    req = requests.get(f"{base_url}{page_url}")
    soup = BeautifulSoup(req.content, "html.parser")
    file_href = soup.find_all(href=re.compile("PEPliste"))
    file_url = file_href[0].get("href")  # do test here for multiple links
    return f"{base_url}{file_url}"


def parse_pep_xlsx(link):
    data = pd.read_excel(link, header=1, index_col=None).reset_index(drop=True)
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
    returned_cols = ["full_name", "type", "source", "address"]

    names_df = data[parsing_cols]
    names_df = names_df.dropna(axis="index")
    names_df = names_df.rename(columns=rename_dict)
    names_df["type"] = TYPE_ENUM[0]
    names_df["source"] = link
    names_df["address"] = None
    names_df["full_name"] = f"{names_df['first_name']} {names_df['sur_name']}"
    names_df = names_df[returned_cols]
    return names_df


def remove_pep_from_db():
    dbtools.delete_all_bad_persons(list_type=TYPE_ENUM[0])


def add_new_pep_to_db(df):
    dbtools.update_df(df, list_type=TYPE_ENUM[0])


def run(event, context):
    # file_link = get_link_to_file()
    # parsed_df = parse_pep_xlsx(file_link)
    # add_new_pep_to_db(parsed_df)
    return {
        "statusCode": 200,
        "body": f"<html><body><p>{os.listdir('.')}</p></body></html>",
        "headers": {"Content-Type": "text/html"},
    }


if __name__ == "__main__":
    file_link = get_link_to_file()
    parsed_df = parse_pep_xlsx(file_link)
    add_new_pep_to_db(parsed_df)
