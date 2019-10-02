
function addressToString (address) {
  return `${address.postnummer} ${address.postdistrikt} ${address.vejnavn} ${(address.husnummerFra || '')} ${(address.bogstavFra || '')}`.trim()
}

function parsePerson (person) {
  return {
    name: person.navne[0].navn,
    countryCode: !!person.beliggenhedsadresse[0]
      ? person.beliggenhedsadresse[0].landekode
      : 'ZZ'
  }
}

function parseMotherCompany (motherCompany) {
  return {
    name: motherCompany.navne[0].navn,
    address: addressToString(motherCompany.beliggenhedsadresse[0]),
    countryCode: motherCompany.beliggenhedsadresse[0].landekode,
    vat: 'DK-' + motherCompany.forretningsnoegle,
    startingDate: motherCompany.sidstOpdateret
  }
}

function parseCompany (company) {
  return {
    name: company.virksomhedMetadata.nyesteNavn.navn,
    address: addressToString(company.virksomhedMetadata.nyesteBeliggenhedsadresse),
    countryCode: company.virksomhedMetadata.nyesteBeliggenhedsadresse.landekode,
    vat: 'DK-' + company.cvrNummer.toString().padStart(8, '0'),
    startingDate: company.virksomhedMetadata.stiftelsesDato
  }
}

function parse (hit) {
  const entry = hit._source.Vrvirksomhed
	const persons = []
  const motherCompanies = []

  const company = parseCompany(entry)
  
	entry.deltagerRelation.forEach(entity => {
    try {
      const hasType = !!entity.deltager.enhedstype
      if (!hasType) return
    }
    catch (error) {
      console.log(entity)
      console.log(error)
      return
    }

    const isPerson = entity.deltager.enhedstype === 'PERSON'
    const isCompany = entity.deltager.enhedstype === 'VIRKSOMHED'

    if (isPerson) persons.push(parsePerson(entity.deltager))
    if (isCompany) motherCompanies.push(parseMotherCompany(entity.deltager))
  })

  company.persons = persons
  company.motherCompanies = motherCompanies

  return company
}

module.exports = { parse }
