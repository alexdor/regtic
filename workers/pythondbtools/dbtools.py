import os
import uuid
import enum
from sqlalchemy import Column, String, create_engine, Enum, ForeignKey
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


class BAD_PERSON_TYPE(enum.Enum):
    PEP = "PEP"
    SANCTION = "SANCTION"


class BadPerson(base):
    __tablename__ = "bad_persons"
    id = Column(UUID(as_uuid=True), primary_key=True, unique=True, default=uuid.uuid4)
    full_name = Column(String)
    type = Column(Enum(BAD_PERSON_TYPE))
    source = Column(String)
    address = Column(String)


class BadPersonAlias(base):
    __tablename__ = "bad_persons_aliases"
    id = Column(UUID(as_uuid=True), primary_key=True, unique=True, default=uuid.uuid4)
    full_name = Column(String)
    bad_person_id = Column(UUID(as_uuid=True), ForeignKey('bad_persons.id'), nullable=False)


def update_df(df, list_type):
    session = Session()

    try:
        delete_all_bad_persons_in_session(session, list_type)

        if list_type == BAD_PERSON_TYPE.SANCTION:
            for index, bad_person in df.iterrows():
                bad_person_obj = BadPerson(
                    full_name=bad_person["full_name"],
                    type=bad_person["type"],
                    source=bad_person["source"],
                    address=bad_person["address"],
                )
                session.add(bad_person_obj)
                session.commit()
                for alias in bad_person["alias"]:
                    alias_obj = BadPersonAlias(
                        full_name=alias,
                        bad_person_id=bad_person_obj.id,
                    )
                    session.add(alias_obj)
            session.commit()
        else:
            for index, bad_person in df.iterrows():
                bad_person_obj = BadPerson(
                    full_name=bad_person["full_name"],
                    type=bad_person["type"],
                    source=bad_person["source"],
                    address=bad_person["address"],
                )
                session.add(bad_person_obj)
            session.commit()

    except Exception as err:
        session.rollback()
        raise err
    finally:
        session.close()


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
        if list_type == BAD_PERSON_TYPE.SANCTION:
            bad_persons_aliases = session.query(BadPersonAlias)
            for bad_persons_alias in bad_persons_aliases:
                session.delete(bad_persons_alias)

        bad_persons = session.query(BadPerson).filter(BadPerson.type == list_type)

    else:
        bad_persons = session.query(BadPerson)

    for bad_person in bad_persons:
        session.delete(bad_person)


def delete_all_bad_persons(list_type=None):  # Not sure about this
    session = Session()

    try:
        if list_type is not None:
            bad_persons = session.query(BadPerson).filter(BadPerson.type == list_type)
            if list_type == BAD_PERSON_TYPE.SANCTION:
                bad_persons_alias = session.query(BadPersonAlias)
        else:
            bad_persons = session.query(BadPerson)
            bad_persons_alias = session.query(BadPersonAlias)

        if list_type == BAD_PERSON_TYPE.SANCTION:
            for bad_person in bad_persons:
                session.delete(bad_person)
                session.delete(bad_persons_alias)

        else:
            for bad_person in bad_persons:
                session.delete(bad_person)

    except Exception as err:
        session.rollback()
        raise err
    finally:
        session.close()

