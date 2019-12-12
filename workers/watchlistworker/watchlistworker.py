from sqlalchemy import create_engine, func
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from workers.pythondbtools.dbtools import Persons, BadPersonAlias
import os

from psycopg2.extensions import AsIs
from django.contrib.postgres.search import SearchRank

db_uri = os.environ["REGTIC_DATABASE_URL"]

base = declarative_base()
db = create_engine(db_uri)
Session = sessionmaker(db)
base.metadata.create_all(db)

created_bad_people = ["a78b77bf-c91c-47f4-b85a-314f7d778038", "4aea249c-c967-46ff-ab41-239e0a786eee", "454ba45f-a350-45c4-80ee-5f92138ff970", "23e3b13d-d26b-40f0-9f04-fc44290542dd", "c031b078-8336-4e83-8a11-202da3e03adc"]
updated_bad_people = ["4aea249c-c967-46ff-ab41-239e0a786eee", "454ba45f-a350-45c4-80ee-5f92138ff970", "23e3b13d-d26b-40f0-9f04-fc44290542dd", "c031b078-8336-4e83-8a11-202da3e03adc"]
deleted_bad_people = ["4aea249c-c967-46ff-ab41-239e0a786eee", "454ba45f-a350-45c4-80ee-5f92138ff970", "23e3b13d-d26b-40f0-9f04-fc44290542dd", "c031b078-8336-4e83-8a11-202da3e03adc"]

watchlist_companies = ["Direktoratet för statens säkerhet [Syrien]", "Direktoratet för säkerhetstjänsten [Syrien]", "ředitelství pro vojenské zpravodajství [Sýrie]", "Serviço de Informações da Força Aérea [Síria]"]

Session = sessionmaker(db)

    # " @@ plainto_tsquery('simple', ? )"
    # check tsrank


def worker(created_bad_people, updated_bad_people, deleted_bad_people, watchlist_companies):

    try:
        session = Session()

        bad_person_name = []

        persons_obj = Persons(
            first_name="Asadollah",
            last_name="JAFARI",
            country_code="ZZ"
        )
        session.add(persons_obj)
        session.commit()

        for person_id in created_bad_people:
            print("Person id: " + person_id)
            bad_person_vector = (session.query(BadPersonAlias.name_vector).filter_by(bad_person_id=person_id).first()[0])
            print("Bad person vector: " + str(bad_person_vector))
            bad_person_split = bad_person_vector.split("'")
            bad_person = []
            for i in range(len(bad_person_split)):
                if i % 2 != 0:
                    bad_person.append(bad_person_split[i])

            print("Bad person: " + str(bad_person))
            bad_person_name.append(bad_person_vector)

            #session.query(Persons).filter("Persons.name_vector @@ to_tsquery(:search_string)").params(search_string=bad_person_vector).all()
            #tsquery = session.query(Persons).filter(Persons.name_vector.op('@@')(func.to_tsquery(bad_person_name[0]))).all()
            #print("Tsquery: " + str(tsquery))

            tsquery_2 = session.query(Persons.name_vector).first()
            aux_query = session.query(bad_person_vector, tsquery_2)

            # search_vector = SearchVector('body_text')
            # search_query = SearchQuery('cheese')
            search_rank = SearchRank(bad_person_vector, tsquery_2)
            # print("QUE VA: " + str(session.execute("SELECT name_vector FROM Persons").first()))
            session.execute("SELECT ts_rank(%s, %s) FROM %s", (AsIs(bad_person_vector), AsIs(tsquery_2), AsIs(aux_query)))

            tryRank = "*, ts_rank(" + bad_person_vector + ", " + tsquery_2 + ") as rank"
            print(tryRank)



            print("Tsquery: " + str(tsquery_2))
            print("Search rank: " + str(search_rank))


            #rank = func.ts_rank(bad_person_vector, tsquery_2)
            #print(rank)

            # query Persons table to obtain the names that match the name vector of persons table
            # select persons from Persons where person_name_name_vector = person_name_name_vector
            # if (bad_person_name.to_tsquery() == persons.nameVector) > 0.7:
            #       notify()
            # session.query(Person).filter("t.column_name @@ to_tsquery(:search_string)").params(search_string=search_string).all()


    except Exception as err:
        session.rollback()
        raise err
    finally:
        session.close()


worker(created_bad_people, updated_bad_people, deleted_bad_people, watchlist_companies)


