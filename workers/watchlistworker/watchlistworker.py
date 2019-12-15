from sqlalchemy import create_engine, func
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from sqlalchemy.sql import text
from workers.pythondbtools import dbtools
import os

db_uri = os.environ["REGTIC_DATABASE_URL"]


base = declarative_base()
db = create_engine(db_uri)
Session = sessionmaker(db)
base.metadata.create_all(db)


def add_to_bad_person_to_person_dict(bp_to_p_dict, bad_person_id, session):
    bad_person_alias_query = (
        session.query(dbtools.BadPersonAllNames.full_name)
        .filter_by(bad_person_id=bad_person_id)
        .all()
    )
    for alias_query in bad_person_alias_query:
        ts_query = (
            session.query(
                dbtools.Persons,
                func.ts_rank(
                    dbtools.Persons.name_vector,
                    func.plainto_tsquery("simple", alias_query[0]),
                ).label("rank"),
            )
            .filter(
                dbtools.Persons.name_vector.op("@@")(
                    func.plainto_tsquery("simple", alias_query[0])
                )
            )
            .order_by(text("rank desc"))
        ).one_or_none()

        if ts_query is not None:
            top_result = ts_query[0]
            bp_to_p_dict[bad_person_id] = top_result.id


def add_to_bad_company_to_company_dict(bc_to_c_dict, bad_company_id, session):
    bad_company_alias_query = (
        session.query(dbtools.BadCompanyAllNames.name)
        .filter_by(bad_company_id=bad_company_id)
        .all()
    )
    for alias_query in bad_company_alias_query:
        ts_query = (
            session.query(
                dbtools.Companies,
                func.ts_rank(
                    dbtools.Companies.name_vector,
                    func.plainto_tsquery("simple", alias_query[0]),
                ).label("rank"),
            )
            .filter(
                dbtools.Companies.name_vector.op("@@")(
                    func.plainto_tsquery("simple", alias_query[0])
                )
            )
            .order_by(text("rank desc"))
        ).one_or_none()

        if ts_query is not None:
            top_result = ts_query[0]
            bc_to_c_dict[bad_company_id] = top_result.id


def inserted_person_worker(created_bad_people):
    session = Session()
    try:
        created_bp_to_p_dict = {}
        for person_id in created_bad_people:
            add_to_bad_person_to_person_dict(created_bp_to_p_dict, person_id, session)

        for key in created_bp_to_p_dict:
            print(key)
            print(created_bp_to_p_dict[key])
            bp_to_p_obj = dbtools.BadPersonToPerson(
                bad_person_id=key, person_id=created_bp_to_p_dict[key]
            )
            session.add(bp_to_p_obj)

        session.commit()

    except Exception as err:
        session.rollback()
        raise err
    finally:
        session.close()


def inserted_company_worker(created_bad_companies):
    session = Session()
    try:
        created_bc_to_c_dict = {}
        for company_id in created_bad_companies:
            add_to_bad_company_to_company_dict(
                created_bc_to_c_dict, company_id, session
            )

        for key in created_bc_to_c_dict:
            print(key)
            print(created_bc_to_c_dict[key])
            bc_to_c_obj = dbtools.BadCompanyToCompany(
                bad_company_id=key, company_id=created_bc_to_c_dict[key]
            )
            session.add(bc_to_c_obj)

        session.commit()

    except Exception as err:
        session.rollback()
        raise err
    finally:
        session.close()
