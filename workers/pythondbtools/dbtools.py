import os
import uuid
import enum
from datetime import datetime
from sqlalchemy import Column, String, create_engine, Enum, cast
from sqlalchemy.dialects.postgresql import UUID, TIMESTAMP, VARCHAR, ARRAY
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
# db_uri = "postgres://admin:admin@localhost:5432/regtic" # todo remove
base = declarative_base()
db = create_engine(db_uri)
Session = sessionmaker(db)
base.metadata.create_all(db)


class BAD_PERSON_TYPE(enum.Enum):
    PEP = "PEP"
    SANCTION = "SANCTION"


class BadPerson(base):
    __tablename__ = "bad_persons"
    id = Column(UUID(as_uuid=True), primary_key=True, unique=True, default=uuid.uuid4)
    full_name = Column(String)
    type = Column(Enum(BAD_PERSON_TYPE))
    source = Column(String)
    updated_at = Column(TIMESTAMP, default=datetime.utcnow, onupdate=datetime.utcnow)
    citizenship_country_code = Column(ARRAY(VARCHAR(2)))


def push_bad_person(full_name, list_type, source, country_code):
    session = Session()
    bad_person = BadPerson(
        full_name=full_name,
        type=list_type,
        source=source,
        citizenship_country_code=country_code,
    )
    try:
        session.add(bad_person)
        session.commit()
    except Exception as err:
        session.rollback()
        raise err
    finally:
        session.close()


def update_df(df, list_type):
    session = Session()

    try:
        delete_all_bad_persons_in_session(session, list_type)

        for index, bad_person in df.iterrows():
            bad_person_obj = BadPerson(
                full_name=bad_person["full_name"],
                type=list_type,
                source=bad_person["source"],
                citizenship_country_code=bad_person["country_code"],
            )
            session.add(bad_person_obj)

        session.commit()
    except Exception as err:
        session.rollback()
        raise err
    finally:
        session.close()


def upsert_df(df, list_type):
    time_now = datetime.utcnow()

    # Get persons from database
    session = Session()

    inserted_ids = []
    updated_ids = []
    deleted_ids = []

    for _, person in df.iterrows():
        query_person = (
            session.query(BadPerson)
            .filter(BadPerson.type == list_type)
            .filter(BadPerson.full_name == person["full_name"])
            .filter(
                BadPerson.citizenship_country_code.contains(
                    cast(person["country_code"], ARRAY(VARCHAR(2)))
                )
            )
            .one_or_none()
        )

        if query_person is None:
            bad_person_obj = BadPerson(
                full_name=person["full_name"],
                type=person["type"],
                source=person["source"],
                citizenship_country_code=person["country_code"],
            )

            session.add(bad_person_obj)
            inserted_ids.append(bad_person_obj.id)
        else:
            updated_ids.append(query_person.id)
            query_person.updated_at = time_now

    session.commit()

    removed_persons = (
        session.query(BadPerson)
        .filter(BadPerson.type == list_type)
        .filter(BadPerson.updated_at < time_now)
        .all()
    )

    for person in removed_persons:
        deleted_ids.append(person.id)

    updates_trigger(updated_ids)
    inserts_trigger(inserted_ids)
    deletes_trigger(deleted_ids)

    return f"inserted: {len(inserted_ids)} updated: {len(updated_ids)} deleted: {len(deleted_ids)}"


def updates_trigger(UUIDS):
    # Do cool stuff
    return 1


def inserts_trigger(UUIDS):
    # Do cool stuff
    return 1


def deletes_trigger(UUIDS):
    # Do cool stuff
    return 1


def read_bad_persons(list_type=None):
    session = Session()
    bad_persons = []

    try:
        if type is not None:
            bad_persons = session.query(BadPerson).filter(BadPerson.type == list_type)
        else:
            bad_persons = session.query(BadPerson)
        session.commit()
    except Exception as err:
        session.rollback()
        raise err
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
    except Exception as err:
        session.rollback()
        raise err
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
    except Exception as err:
        session.rollback()
        raise err
    finally:
        session.close()
