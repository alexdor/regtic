import os
import uuid
from enum import Enum

import yaml
from sqlalchemy import Column, String, create_engine, inspect
from sqlalchemy.dialects.postgresql import UUID
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

"""
id              (UUID)
full_name       (None)
type            (None)
source          (None)
address         (None)
updated_at      (Now)
created_at      (Now)
"""

db_uri = os.environ["REGTIC_DATABASE_URL"]

base = declarative_base()
db = create_engine(db_uri)
Session = sessionmaker(db)
base.metadata.create_all(db)


class BAD_PERSON_TYPE(Enum):
    PEP = "PEP"
    SANCTION = "SANCTION"


class BadPerson(base):
    __tablename__ = "bad_persons"
    id = Column(UUID(as_uuid=True), primary_key=True, unique=True, default=uuid.uuid4)
    full_name = Column(String)
    type = Column(String)
    source = Column(String)
    address = Column(String)


def push_bad_person(full_name, list_type, source, address):
    session = Session()
    bad_person = BadPerson(
        full_name=full_name, type=list_type, source=source, address=address
    )
    try:
        session.add(bad_person)
        session.commit()
    except:
        session.rollback()
        raise
    finally:
        session.close()


def update_df(df, list_type):
    session = Session()

    try:
        delete_all_bad_persons_in_session(session, list_type)

        for index, bad_person in df.iterrows():
            bad_person_obj = BadPerson(
                full_name=bad_person["full_name"],
                type=bad_person["type"],
                source=bad_person["source"],
                address=bad_person["address"],
            )
            session.add(bad_person_obj)

        session.commit()
    except:
        session.rollback()
    finally:
        session.close()


def read_bad_persons(list_type=None):
    session = Session()
    bad_persons = []

    try:
        if type is not None:
            bad_persons = session.query(BadPerson).filter(BadPerson.type == list_type)
        else:
            bad_persons = session.query(BadPerson)
        session.commit()
    except:
        session.rollback()
    finally:
        session.close()

    bad_persons_list = []
    for bad_person in bad_persons:
        bad_persons_list.append(bad_person)
    return bad_persons_list


def delete_bad_person(bad_person):
    session = Session()
    try:
        session.delete(bad_person)
        session.commit()
    except:
        session.rollback()
    finally:
        session.close()


def delete_all_bad_persons_in_session(session=None, list_type=None):
    if list_type is not None:
        bad_persons = session.query(BadPerson).filter(BadPerson.type == list_type)
    else:
        bad_persons = session.query(BadPerson)

    for bad_person in bad_persons:
        session.delete(bad_person)


def delete_all_bad_persons(list_type=None):
    session = Session()

    try:
        if list_type is not None:
            bad_persons = session.query(BadPerson).filter(BadPerson.type == list_type)
        else:
            bad_persons = session.query(BadPerson)

        for bad_person in bad_persons:
            session.delete(bad_person)
    except:
        session.rollback()
    finally:
        session.close()
