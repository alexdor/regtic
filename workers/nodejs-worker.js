
'use strict'

module.exports.hello = async (event, context) => {
  return { statusCode: 200, body: 'Hello world!' }
}
