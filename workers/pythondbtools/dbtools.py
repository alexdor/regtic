from sqlalchemy import create_engine
from sqlalchemy import Column, String
from sqlalchemy.dialects.postgresql import UUID
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import inspect
from sqlalchemy.orm import sessionmaker
import uuid

"""
id              (UUID)
full_name       (None)
type            (None)
source          (None)
address         (None)
updated_at      (Now)
created_at      (Now)
"""

db_uri = "postgres://admin:admin@localhost:5432/regtic"
base = declarative_base()
db = create_engine(db_uri)
Session = sessionmaker(db)
session = Session()
base.metadata.create_all(db)


class BadPerson(base):
    __tablename__ = "bad_persons"
    id = Column(UUID(as_uuid=True), primary_key=True, unique=True, default=uuid.uuid4)
    full_name = Column(String)
    type = Column(String)
    source = Column(String)
    address = Column(String)


def push_bad_person(session, full_name, type, source, address):
    bad_person = BadPerson(
        full_name=full_name, type=type, source=source, address=address
    )
    session.add(bad_person)
    session.commit()


def push_df(df):
    for index, bad_person in df.iterrows():
        bad_person_obj = BadPerson(
            full_name=bad_person["full_name"],
            type=bad_person["type"],
            source=bad_person["source"],
            address=bad_person["address"],
        )
        session.add(bad_person_obj)
    session.commit()


def read_bad_persons():
    bad_persons = session.query(BadPerson)
    bad_persons_list = []
    for bad_person in bad_persons:
        bad_persons_list.append(bad_person)
    return bad_persons_list


def delete_bad_person(bad_person):
    session.delete(bad_person)
    session.commit()


def delete_all_bad_persons():
    bad_persons = session.query(BadPerson)
    for bad_person in bad_persons:
        session.delete(bad_person)
        session.commit()


if __name__ == "__main__":
    engine = create_engine(db_uri)
    inspector = inspect(engine)
    print(inspector.get_table_names())
    print(inspector.get_columns("bad_persons"))
    delete_all_bad_persons()
