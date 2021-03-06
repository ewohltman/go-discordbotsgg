package mock

const botResponseString = `
{
  "userId": "botID",
  "clientId": "clientID",
  "username": "Test Bot 1",
  "discriminator": null,
  "avatarURL": null,
  "coOwners": [],
  "prefix": "testBot",
  "helpCommand": "testBot",
  "libraryName": "discordgo",
  "website": null,
  "supportInvite": null,
  "botInvite": null,
  "shortDescription": null,
  "longDescription": null,
  "openSource": null,
  "shardCount": 1,
  "guildCount": 0,
  "verified": true,
  "online": true,
  "inGuild": true,
  "deleted": false,
  "owner": {
    "username": "testOwner",
    "discriminator": null,
    "userId": "112358"
  },
  "addedDate": "2016-10-30T04:59:04.000Z",
  "status": "online"
}
`

const botsResponseString = `
{
  "count": 2,
  "limit": 50,
  "page": 0,
  "bots": [
    {
      "userId": "botID",
      "clientId": "clientID",
      "username": "Test Bot 1",
      "discriminator": null,
      "avatarURL": null,
      "coOwners": [],
      "prefix": "testBot",
      "helpCommand": "testBot",
      "libraryName": "discordgo",
      "website": null,
      "supportInvite": null,
      "botInvite": null,
      "shortDescription": null,
      "longDescription": null,
      "openSource": null,
      "shardCount": 1,
      "guildCount": 0,
      "verified": true,
      "online": true,
      "inGuild": true,
      "deleted": false,
      "owner": {
        "username": "testOwner",
        "discriminator": null,
        "userId": "112358"
      },
      "addedDate": "2016-10-30T04:59:04.000Z",
      "status": "online"
    },
    {
      "userId": "12345",
      "clientId": "67890",
      "username": "Test Bot 2",
      "discriminator": null,
      "avatarURL": null,
      "coOwners": [],
      "prefix": "testBot",
      "helpCommand": "testBot",
      "libraryName": "discordgo",
      "website": null,
      "supportInvite": null,
      "botInvite": null,
      "shortDescription": null,
      "longDescription": null,
      "openSource": null,
      "shardCount": 1,
      "guildCount": 0,
      "verified": true,
      "online": true,
      "inGuild": true,
      "deleted": false,
      "owner": {
        "username": "testOwner",
        "discriminator": null,
        "userId": "112358"
      },
      "addedDate": "2016-10-30T04:59:04.000Z",
      "status": "online"
    }
  ]
}
`
