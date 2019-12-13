function addressToString(address) {
  return `${address.postnummer} ${address.postdistrikt} ${
    address.vejnavn
  } ${address.husnummerFra || ""} ${address.bogstavFra || ""}`.trim();
}

function getName(entity) {
  const entityNameObj = entity.navne.find(
    navnObj => !navnObj.periode.gyldigTil
  );
  return entityNameObj ? entityNameObj.navn : null;
}

function getCompanyName(company) {
  return company.virksomhedMetadata.nyesteNavn
    ? company.virksomhedMetadata.nyesteNavn.navn
    : company.navne.slice(-1)[0];
}

function getCountryCode(entity) {
  const entityLocationObj = entity.beliggenhedsadresse.find(
    ba => !ba.periode.gyldigTil
  );
  return entityLocationObj ? entityLocationObj.landekode : "ZZ";
}

function parsePerson(person) {
  return {
    firstName: getName(person)
      .split(" ")
      .slice(0, -1)
      .join(" "),
    lastName: getName(person)
      .split(" ")
      .slice(-1)
      .join(" "),
    countryCode: getCountryCode(person)
  };
}

function parseMotherCompany(motherCompany) {
  return {
    name: getName(motherCompany),
    address: addressToString(motherCompany.beliggenhedsadresse[0]),
    countryCode: getCountryCode(motherCompany),
    vat: "DK-" + motherCompany.forretningsnoegle,
    startingDate: motherCompany.sidstOpdateret
  };
}

function parseCompany(company) {
  return {
    name: getCompanyName(company),
    address: addressToString(
      company.virksomhedMetadata.nyesteBeliggenhedsadresse
    ),
    countryCode: company.virksomhedMetadata.nyesteBeliggenhedsadresse.landekode,
    vat: "DK-" + company.cvrNummer.toString().padStart(8, "0"),
    startingDate: company.virksomhedMetadata.stiftelsesDato,
    type: company.virksomhedMetadata.nyesteVirksomhedsform.kortBeskrivelse
  };
}

function parseMetaData(titles) {
  function translateTitle(title) {
    const titleMapping = {
      bestyrelse: "board of directors",
      direktion: "direction",
      ejerregister: "legal owner",
      "reelle ejere": "ultimate beneficial owner",
      revision: "accountant",
      stiftere: "founder"
    };

    title = title.toLowerCase();

    return titleMapping[title];
  }

  const relations = titles.reduce((result, relation) => {
    const isActiveRelation =
      relation.organisationsNavn[0].periode.gyldigTil === null;
    if (!isActiveRelation) return result;

    const translatedRelation = translateTitle(
      relation.organisationsNavn[0].navn
    );

    const isTranslatedRelationValid = !!translatedRelation;
    if (!isTranslatedRelationValid) return result;

    result.push(translatedRelation);

    return result;
  }, []);

  return { relations };
}

function parse(hit) {
  const entry = hit._source.Vrvirksomhed;
  const persons = [];
  const motherCompanies = [];

  const company = parseCompany(entry);

  const isEntryValid = !!entry.deltagerRelation;
  if (!isEntryValid) return;

  entry.deltagerRelation.forEach(entity => {
    const hasDeltager = !!entity.deltager;
    const hasType = hasDeltager && entity.deltager.enhedstype;
    if (!hasType) return;

    const metaData = parseMetaData(entity.organisationer);

    const isPerson = entity.deltager.enhedstype === "PERSON";
    if (isPerson) {
      persons.push({
        ...parsePerson(entity.deltager),
        ...metaData
      });
    }

    const isCompany = entity.deltager.enhedstype === "VIRKSOMHED";
    if (isCompany) {
      motherCompanies.push({
        ...parseMotherCompany(entity.deltager),
        ...metaData
      });
    }
  });

  company.persons = persons;
  company.motherCompanies = motherCompanies;

  return company;
}

module.exports = { parse };
