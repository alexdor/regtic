import os
import uuid
import enum
from sqlalchemy import Column, String, create_engine, Enum, ForeignKey
from sqlalchemy.dialects.postgresql import UUID, VARCHAR, ARRAY
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

db_uri = os.environ["REGTIC_DATABASE_URL"]
# db_uri = "postgres://admin:admin@localhost:5432/regtic"

base = declarative_base()
db = create_engine(db_uri)
Session = sessionmaker(db)
base.metadata.create_all(db)


class BAD_PERSON_TYPE(enum.Enum):
    PEP = "PEP"
    SANCTION = "SANCTION"


class ADDRESS_TYPE(enum.Enum):
    BIRTHPLACE = "birthplace"
    ADDRESS = "address"


class BadPerson(base):
    __tablename__ = "bad_persons"
    id = Column(UUID(as_uuid=True), primary_key=True, unique=True, default=uuid.uuid4)
    full_name = Column(String)
    type = Column(Enum(BAD_PERSON_TYPE))
    source = Column(String)
    address = Column(String)
    citizenship_country_code = Column(ARRAY(VARCHAR(2)))


class BadPersonAlias(base):
    __tablename__ = "bad_persons_aliases"
    id = Column(UUID(as_uuid=True), primary_key=True, unique=True, default=uuid.uuid4)
    full_name = Column(String)
    bad_person_id = Column(UUID(as_uuid=True), ForeignKey('bad_persons.id'), nullable=False)


class BadCompany(base):
    __tablename__ = "bad_companies"
    id = Column(UUID(as_uuid=True), primary_key=True, unique=True, default=uuid.uuid4)
    name = Column(String)
    address = Column(String)
    source = Column(String)
    citizenship_region = Column(String)
    citizenship_country_code = Column(ARRAY(VARCHAR(2)))
    type = Column(Enum(BAD_PERSON_TYPE))


class BadCompanyAlias(base):
    __tablename__ = "bad_companies_aliases"
    id = Column(UUID(as_uuid=True), primary_key=True, unique=True, default=uuid.uuid4)
    name = Column(String)
    type = Column(Enum(BAD_PERSON_TYPE))
    bad_company_id = Column(UUID(as_uuid=True), ForeignKey('bad_companies.id'))


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
    bad_company_id = Column(UUID(as_uuid=True), ForeignKey('bad_companies.id'), nullable=False)


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
    bad_person_id = Column(UUID(as_uuid=True), ForeignKey('bad_persons.id'), nullable=False)


def update_df(df, list_type):
    session = Session()

    """
    sanctions_df = pd.DataFrame(
        {'address_type': address_type,
         'citizenship_code': citizenship_code,
         'street': street,
         'poBox': poBox,
         'city': city,
         'zipCode': zipCode,
         'region': region,
         'place': place,
         'entity': subject_type,
         'country': country,
         'type': type,
         'full_name': wholeName,
         'alias': alias,
         'source': source
         })
    """

    try:
        delete_all_bad_persons_in_session(session, list_type)

        if list_type == BAD_PERSON_TYPE.SANCTION:

            for index, bad_person in df.iterrows():

                if bad_person["entity"] == "E":
                    bad_company_obj = BadCompany(
                        name=bad_person["full_name"],
                        # address=bad_person["address"],
                        source=bad_person["source"],
                        citizenship_region=bad_person["citizenship_region"],
                        citizenship_country_code=bad_person["citizenship_code"],
                        type=bad_person["type"]
                    )
                    session.add(bad_company_obj)
                    session.commit()
                    for alias in bad_person["alias"]:
                        if session.query(BadCompanyAlias).filter_by(name=alias).first() is None:
                            alias_obj = BadCompanyAlias(
                                name=alias,
                                type=bad_company_obj.type,
                                bad_company_id=bad_company_obj.id
                            )
                            session.add(alias_obj)
                    session.commit()
                    for i in range(len(bad_person["street"])):
                        if session.query(BadCompanyAddresses).filter_by(street=bad_person["street"][i], city=bad_person["city"][i], zip_code=bad_person["zipCode"][i], region=bad_person["region"][i], place=bad_person["place"][i], po_box=bad_person["poBox"][i], country_code=bad_person["country"][i], bad_company_id=bad_company_obj.id).first() is None:
                            bad_company_address_obj = BadCompanyAddresses(
                                street=bad_person["street"][i],
                                city=bad_person["city"][i],
                                zip_code=bad_person["zipCode"][i],
                                region=bad_person["region"][i],
                                place=bad_person["place"][i],
                                po_box=bad_person["poBox"][i],
                                country_code=bad_person["country"][i],
                                bad_company_id=bad_company_obj.id
                            )
                            session.add(bad_company_address_obj)
                    session.commit()

                if bad_person["entity"] == "P":
                    bad_person_obj = BadPerson(
                        full_name=bad_person["full_name"],
                        type=bad_person["type"],
                        source=bad_person["source"],
                        #address=bad_person["address"],
                        citizenship_country_code=bad_person["citizenship_code"]
                    )
                    session.add(bad_person_obj)
                    session.commit()

                    for alias in bad_person["alias"]:
                        if session.query(BadPersonAlias).filter_by(full_name=alias).first() is None:
                            alias_obj = BadPersonAlias(
                                full_name=alias,
                                bad_person_id=bad_person_obj.id,
                            )
                            session.add(alias_obj)
                    session.commit()

                    for i in range(len(bad_person["street"])):
                        if session.query(BadPersonAddresses).filter_by(street=bad_person["street"][i], city=bad_person["city"][i], zip_code=bad_person["zipCode"][i], region=bad_person["region"][i], place=bad_person["place"][i], po_box=bad_person["poBox"][i], country_code=bad_person["country"][i], type=bad_person["address_type"][i].value, bad_person_id=bad_person_obj.id) is None:
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
                    session.commit()

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
            bad_companies = session.query(BadCompany)
            bad_companies_alias = session.query(BadCompanyAlias)
            bad_persons_addresses = session.query(BadPersonAddresses)
            bad_company_addresses = session.query(BadCompanyAddresses)
            bad_persons_aliases = session.query(BadPersonAlias)

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

