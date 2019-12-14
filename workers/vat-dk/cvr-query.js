/* eslint-disable @typescript-eslint/camelcase */

module.exports = {
  size: 200,
  _source: [
    "Vrvirksomhed.cvrNummer",
    "Vrvirksomhed.navne",
    "Vrvirksomhed.virksomhedMetadata",
    "Vrvirksomhed.deltagerRelation"
  ],
  query: {
    bool: {
      must: [{ exists: { field: "Vrvirksomhed.cvrNummer" } }],
      must_not: [
        { exists: { field: "Vrvirksomhed.livsforloeb.periode.gyldigTil" } }
      ]
    }
  }
};
