import enum
import os
import uuid
from datetime import datetime
from sqlalchemy import Column, String, create_engine, Enum, cast, ForeignKey
from sqlalchemy.dialects.postgresql import UUID, TIMESTAMP, VARCHAR, ARRAY
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

db_uri = os.environ["REGTIC_DATABASE_URL"]

base = declarative_base()
db = create_engine(db_uri)
Session = sessionmaker(db)
base.metadata.create_all(db)


class BAD_PERSON_TYPE(enum.Enum):
    PEP = "PEP"
    SANCTION = "SANCTION"


class ADDRESS_TYPE(enum.Enum):
    birthplace = "birthplace"
    address = "address"


class BadPerson(base):
    __tablename__ = "bad_persons"
    id = Column(UUID(as_uuid=True), primary_key=True, unique=True, default=uuid.uuid4)
    full_name = Column(String)
    type = Column(Enum(BAD_PERSON_TYPE))
    source = Column(String)
    updated_at = Column(TIMESTAMP, default=datetime.utcnow, onupdate=datetime.utcnow)
    citizenship_country_code = Column(ARRAY(VARCHAR(2)))


class BadPersonAllNames(base):
    __tablename__ = "bad_persons_all_names"
    id = Column(UUID(as_uuid=True), primary_key=True, unique=True, default=uuid.uuid4)
    full_name = Column(String)
    bad_person_id = Column(
        UUID(as_uuid=True), ForeignKey("bad_persons.id"), nullable=False
    )


class BadCompany(base):
    __tablename__ = "bad_companies"
    id = Column(UUID(as_uuid=True), primary_key=True, unique=True, default=uuid.uuid4)
    name = Column(String)
    address = Column(String)
    source = Column(String)
    citizenship_region = Column(String)
    citizenship_country_code = Column(ARRAY(VARCHAR(2)))
    updated_at = Column(TIMESTAMP, default=datetime.utcnow, onupdate=datetime.utcnow)
    type = Column(Enum(BAD_PERSON_TYPE))


class BadCompanyAllNames(base):
    __tablename__ = "bad_companies_all_names"
    id = Column(UUID(as_uuid=True), primary_key=True, unique=True, default=uuid.uuid4)
    name = Column(String)
    bad_company_id = Column(UUID(as_uuid=True), ForeignKey("bad_companies.id"))


class BadCompanyAddresses(base):
    __tablename__ = "bad_companies_addresses"
    id = Column(UUID(as_uuid=True), primary_key=True, unique=True, default=uuid.uuid4)
    street = Column(String)
    city = Column(String)
    zip_code = Column(String)
    region = Column(String)
    place = Column(String)
    po_box = Column(String)
    country_code = Column(VARCHAR(2))
    bad_company_id = Column(
        UUID(as_uuid=True), ForeignKey("bad_companies.id"), nullable=False
    )


class BadPersonAddresses(base):
    __tablename__ = "bad_persons_addresses"
    id = Column(UUID(as_uuid=True), primary_key=True, unique=True, default=uuid.uuid4)
    street = Column(String)
    city = Column(String)
    zip_code = Column(String)
    region = Column(String)
    place = Column(String)
    po_box = Column(String)
    country_code = Column(VARCHAR(2))
    type = Column(Enum(ADDRESS_TYPE))
    bad_person_id = Column(
        UUID(as_uuid=True), ForeignKey("bad_persons.id"), nullable=False
    )


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
            push_person_in_session(
                session=session, bad_person=bad_person, list_type=list_type
            )
        session.commit()
    except Exception as err:
        session.rollback()
        raise err
    finally:
        session.close()


def push_person_in_session(session, bad_person, list_type):
    returned_id = None
    returned_id_type = "P"

    if list_type == BAD_PERSON_TYPE.PEP:
        bad_person_obj = BadPerson(
            full_name=bad_person["full_name"],
            type=bad_person["type"],
            source=bad_person["source"],
            citizenship_country_code=bad_person["country_code"],
        )
        session.add(bad_person_obj)
        if (
            session.query(BadPersonAllNames)
            .filter_by(full_name=bad_person["full_name"])
            .first()
            is None
        ):
            alias_obj = BadPersonAllNames(
                full_name=bad_person["full_name"], bad_person_id=bad_person_obj.id
            )
            session.add(alias_obj)

        returned_id = bad_person_obj.id

    elif list_type == BAD_PERSON_TYPE.SANCTION:
        if bad_person["entity"] == "E":
            bad_company_obj = BadCompany(
                name=bad_person["full_name"],
                source=bad_person["source"],
                citizenship_region=bad_person["citizenship_region"],
                citizenship_country_code=bad_person["citizenship_code"],
                type=bad_person["type"],
            )
            session.add(bad_company_obj)
            returned_id = bad_company_obj.id
            returned_id_type = "E"
            print("test")
            for alias in bad_person["alias"]:
                if (
                    session.query(BadCompanyAllNames).filter_by(name=alias).first()
                    is None
                ):
                    alias_obj = BadCompanyAllNames(
                        name=alias, bad_company_id=bad_company_obj.id
                    )
                    session.add(alias_obj)

            for i in range(len(bad_person["street"])):
                if (
                    session.query(BadCompanyAddresses)
                    .filter_by(
                        street=bad_person["street"][i],
                        city=bad_person["city"][i],
                        zip_code=bad_person["zipCode"][i],
                        region=bad_person["region"][i],
                        place=bad_person["place"][i],
                        po_box=bad_person["poBox"][i],
                        country_code=bad_person["country"][i],
                        bad_company_id=bad_company_obj.id,
                    )
                    .first()
                    is None
                ):
                    bad_company_address_obj = BadCompanyAddresses(
                        street=bad_person["street"][i],
                        city=bad_person["city"][i],
                        zip_code=bad_person["zipCode"][i],
                        region=bad_person["region"][i],
                        place=bad_person["place"][i],
                        po_box=bad_person["poBox"][i],
                        country_code=bad_person["country"][i],
                        bad_company_id=bad_company_obj.id,
                    )
                    session.add(bad_company_address_obj)

        if bad_person["entity"] == "P":
            bad_person_obj = BadPerson(
                full_name=bad_person["full_name"],
                type=bad_person["type"],
                source=bad_person["source"],
                citizenship_country_code=bad_person["citizenship_code"],
            )
            session.add(bad_person_obj)
            returned_id = bad_person_obj.id

            for alias in bad_person["alias"]:
                if (
                    session.query(BadPersonAllNames).filter_by(full_name=alias).first()
                    is None
                ):
                    alias_obj = BadPersonAllNames(
                        full_name=alias, bad_person_id=bad_person_obj.id
                    )
                    session.add(alias_obj)

            for i in range(len(bad_person["street"])):
                if (
                    session.query(BadPersonAddresses)
                    .filter_by(
                        street=bad_person["street"][i],
                        city=bad_person["city"][i],
                        zip_code=bad_person["zipCode"][i],
                        region=bad_person["region"][i],
                        place=bad_person["place"][i],
                        po_box=bad_person["poBox"][i],
                        country_code=bad_person["country"][i],
                        bad_person_id=bad_person_obj.id,
                    )
                    .first()
                    is None
                ):
                    bad_person_address_obj = BadPersonAddresses(
                        street=bad_person["street"][i],
                        city=bad_person["city"][i],
                        zip_code=bad_person["zipCode"][i],
                        region=bad_person["region"][i],
                        place=bad_person["place"][i],
                        po_box=bad_person["poBox"][i],
                        country_code=bad_person["country"][i],
                        type=bad_person["address_type"][i].value,
                        bad_person_id=bad_person_obj.id,
                    )
                    session.add(bad_person_address_obj)

    return returned_id, returned_id_type


def upsert_df(df, list_type):
    time_now = datetime.utcnow()

    # Get persons from database
    session = Session()

    try:
        inserted_ids = {"P": [], "E": []}
        updated_ids = {"P": [], "E": []}
        deleted_ids = {"P": [], "E": []}

        for _, row in df.iterrows():
            row_type_switch = "P"
            query_row = None
            if list_type == BAD_PERSON_TYPE.SANCTION and row["entity"] == "E":
                query_row = (
                    session.query(BadCompany)
                    .filter(BadCompany.type == list_type)
                    .filter(BadCompany.name == row["full_name"])
                    .one_or_none()
                )
                row_type_switch = "E"
            else:
                test = "true"
                query_row = (
                    session.query(BadPerson)
                    .filter(BadPerson.type == list_type)
                    .filter(BadPerson.full_name == row["full_name"])
                    .one_or_none()
                )

            if query_row is None:
                person_id = push_person_in_session(
                    session=session, bad_person=row, list_type=list_type
                )
                inserted_ids[row_type_switch].append(person_id)
            else:
                updated_ids[row_type_switch].append(query_row.id)
                query_row.updated_at = time_now

        session.commit()

        removed_persons = (
            session.query(BadPerson)
            .filter(BadPerson.type == list_type)
            .filter(BadPerson.updated_at < time_now)
            .all()
        )

        removed_companies = (
            session.query(BadCompany)
            .filter(BadCompany.type == list_type)
            .filter(BadCompany.updated_at < time_now)
            .all()
        )

        for person in removed_persons:
            deleted_ids["P"].append(person.id)

        for company in removed_companies:
            deleted_ids["E"].append(company.id)

        updates_trigger(updated_ids)
        inserts_trigger(inserted_ids)
        deletes_trigger(deleted_ids)

        return (
            f"PERSONS: inserted: {len(inserted_ids['P'])} updated: {len(updated_ids['P'])} deleted: {len(deleted_ids['P'])}\n"
            f"COMPANIES: inserted: {len(inserted_ids['E'])} updated: {len(updated_ids['E'])} deleted: {len(deleted_ids['E'])}\n"
        )
    except Exception as err:
        session.rollback()
        raise err
    finally:
        session.close()


def updates_trigger(uuids):
    # Do cool stuff
    return 1


def inserts_trigger(uuids):
    # Do cool stuff
    return 1


def deletes_trigger(uuids):
    # Do cool stuff
    return 1


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
            bad_companies = session.query(BadCompany)
            bad_companies_alias = session.query(BadCompanyAllNames)
            bad_persons_addresses = session.query(BadPersonAddresses)
            bad_company_addresses = session.query(BadCompanyAddresses)
            bad_persons_aliases = session.query(BadPersonAllNames)

            for bad_company_alias in bad_companies_alias:
                session.delete(bad_company_alias)

            for bad_company_address in bad_company_addresses:
                session.delete(bad_company_address)

            for bad_persons_alias in bad_persons_aliases:
                session.delete(bad_persons_alias)

            for bad_person_address in bad_persons_addresses:
                session.delete(bad_person_address)

            for bad_company in bad_companies:
                session.delete(bad_company)

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
            if list_type == BAD_PERSON_TYPE.SANCTION:
                bad_persons_alias = session.query(BadPersonAllNames)
        else:
            bad_persons = session.query(BadPerson)
            bad_persons_alias = session.query(BadPersonAllNames)

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
