from sqlalchemy import create_engine, func
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from sqlalchemy.sql import text
from workers.pythondbtools.dbtools import Persons, BadPersonAlias
from sqlalchemy import select

from psycopg2.extensions import AsIs


# db_uri = os.environ["REGTIC_DATABASE_URL"]
db_uri = (
    "postgres://admin:admin@localhost:5432/regtic"
)  # todo remove this one when we move it to serverless

base = declarative_base()
db = create_engine(db_uri)
Session = sessionmaker(db)
base.metadata.create_all(db)


# todo remove this test stuff
created_bad_people = ["a1d8b9df-b0fd-432b-af98-e90767953499"]
updated_bad_people = ["647dc6da-0ce9-4346-bc16-278a1664ceea"]


# created_bad_people = ["a78b77bf-c91c-47f4-b85a-314f7d778038", "4aea249c-c967-46ff-ab41-239e0a786eee", "454ba45f-a350-45c4-80ee-5f92138ff970", "23e3b13d-d26b-40f0-9f04-fc44290542dd", "c031b078-8336-4e83-8a11-202da3e03adc"]
# updated_bad_people = [
#    "4aea249c-c967-46ff-ab41-239e0a786eee",
#    "454ba45f-a350-45c4-80ee-5f92138ff970",
#    "23e3b13d-d26b-40f0-9f04-fc44290542dd",
#    "c031b078-8336-4e83-8a11-202da3e03adc",
# ]
deleted_bad_people = [
    "4aea249c-c967-46ff-ab41-239e0a786eee",
    "454ba45f-a350-45c4-80ee-5f92138ff970",
    "23e3b13d-d26b-40f0-9f04-fc44290542dd",
    "c031b078-8336-4e83-8a11-202da3e03adc",
]

watchlist_companies = [
    "Direktoratet för statens säkerhet [Syrien]",
    "Direktoratet för säkerhetstjänsten [Syrien]",
    "ředitelství pro vojenské zpravodajství [Sýrie]",
    "Serviço de Informações da Força Aérea [Síria]",
]


def add_to_bad_person_to_person_dict(bp_to_p_dict, bad_person_id, session):
    bad_person_alias_query = (
        session.query(BadPersonAlias.full_name)
        .filter_by(bad_person_id=bad_person_id)
        .all()
    )
    for alias_query in bad_person_alias_query:
        ts_query = (
            session.query(
                Persons,
                func.ts_rank(
                    Persons.name_vector, func.plainto_tsquery("simple", alias_query[0])
                ).label("rank"),
            )
            .filter(
                Persons.name_vector.op("@@")(
                    func.plainto_tsquery("simple", alias_query[0])
                )
            )
            .order_by(text("rank desc"))
        ).one_or_none()

        if ts_query is not None:
            top_result = ts_query[0]
            print(f"Alias: {alias_query[0]}")
            print(f"Person: {top_result.full_name}")
            print(f"rank: {ts_query[1]}")
            bp_to_p_dict[bad_person_id] = top_result.id


def worker(
    created_bad_people, updated_bad_people, deleted_bad_people, watchlist_companies
):

    try:
        session = Session()

        # persons_obj = Persons(first_name="Abu", last_name="Ali", country_code="IQ")
        # session.add(persons_obj)
        # persons_obj2 = Persons(first_name="Lars", last_name="Løkke Rasmussen", country_code="DK")
        # session.add(persons_obj2)

        # session.commit()

        created_bp_to_p_dict = {}
        for person_id in created_bad_people:
            add_to_bad_person_to_person_dict(created_bp_to_p_dict, person_id, session)

        print(created_bp_to_p_dict)

        updated_bp_to_p_dict = {}
        for person_id in updated_bad_people:
            add_to_bad_person_to_person_dict(updated_bp_to_p_dict, person_id, session)

        deleted_bp_to_p_dict = {}
        for person_id in deleted_bad_people:
            add_to_bad_person_to_person_dict(deleted_bp_to_p_dict, person_id, session)

        # todo Add, update and delete all bp-to-p pairs from table

        # todo Do the same for companies

    except Exception as err:
        session.rollback()
        raise err
    finally:
        session.close()


worker(created_bad_people, updated_bad_people, deleted_bad_people, watchlist_companies)
