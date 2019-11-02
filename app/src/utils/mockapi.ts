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
  },
  async checkCompany(id: string) {
    await sleep(3000);
    return {
      name: "Regtic Demo",
      address: "Demo Street 1",
      vat: "DK-12345678",
      companies: [
        {
          id: "7890-56789-01234-5678",
          found: true,
          name: "Regtic Test Name",
          vat: "DK-78901234",
          address: "DTU Lyngby 101 A",
          country_code: "DK",
          starting_date: Date.now()
        },
        {
          id: "6789-56789-01234-5678",
          found: true,
          name: "Listed Regtic",
          vat: "DK-67890123",
          address: "Demo Blv. 5",
          country_code: "DK",
          starting_date: Date.now()
        },
        {
          id: "5678-56789-01234-5678",
          found: false,
          name: "Regtic Law",
          vat: "DK-56789012",
          address: "Test Street 100",
          country_code: "ZZ",
          starting_date: Date.now()
        }
      ],
      people: {
        bad: [
          {
            name: "John Doe",
            source:
              "EU Sanctions List 2019 (www.example.org/path-to-file) - 23/10/2019",
            company_name: "Regtic Test Name",
            company_id: "7890-56789-01234-5678"
          }
        ],
        warning: [
          {
            name: "Jane Doe",
            source:
              "EU PEP List 2019 (www.example.org/path-to-file) - 22/10/2019",
            company_name: "Listed Regtic",
            company_id: "6789-56789-01234-5678"
          }
        ],
        good: [
          {
            name: "James Doe",
            company_name: "Regtic Test Name",
            company_id: "7890-56789-01234-5678"
          },
          {
            name: "Bob Bobbery",
            company_name: "Regtic Test Name",
            company_id: "6789-56789-01234-567"
          },
          {
            name: "Mallory Malware",
            company_name: "Regtic Test Name",
            company_id: "6789-56789-01234-567"
          }
        ]
      }
    };
  }
};
