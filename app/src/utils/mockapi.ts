function sleep(ms: number) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

export default {
  async findCompanies(searchStr: string) {
    await sleep(2000);
    return [
      {
        name: "Demo Regtic",
        id: "1234-56789-01234-5678",
        address: "Demo Street 1",
        vat: "DK-12345678"
      },
      {
        name: "Regtic Example",
        id: "2345-56789-01234-5678",
        address: "Demo Street 1",
        vat: "DK-23456789"
      },
      {
        name: "Regtic A/S",
        id: "3456-56789-01234-5678",
        address: "Demo Street 2",
        vat: "DK-34567890"
      },
      {
        name: "Regtic Finance",
        id: "4567-56789-01234-5678",
        address: "Test Street 100",
        vat: "DK-45678901"
      },
      {
        name: "Regtic Law",
        id: "5678-56789-01234-5678",
        address: "Test Street 100",
        vat: "DK-56789012"
      },
      {
        name: "Listed Regtic",
        id: "6789-56789-01234-5678",
        address: "Demo Blv. 5",
        vat: "DK-67890123"
      },
      {
        name: "Regtic Test Name",
        id: "7890-56789-01234-5678",
        address: "DTU Lyngby 101 A",
        vat: "DK-78901234"
      }
    ];
  }
};
