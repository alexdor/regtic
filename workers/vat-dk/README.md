# VAT crawler for Denmark (CVR registry)

Last updated 16/9/2019.

### Summary

- Elasticsearch v6.2.4 (to check the version [click here](http://distribution.virk.dk/))
- Usual search endpoint [http://distribution.virk.dk/cvr-permanent/](http://distribution.virk.dk/cvr-permanent/)
- Credentials are needed, only danish companies with a valid VAT number can apply for one
- Queries where `from + size >= 3000` needs to be scrolled through

### Gotchas

- When `scroll`ing the initial request needs to be sent here `http://distribution.virk.dk/cvr-permanent/`, all subsequent `scroll` requests needs to be sent here `http://distribution.virk.dk/_search/scroll`
- The documentation sucks
