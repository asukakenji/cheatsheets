# Hyperledger Composer Model File (cto) Cheatsheet

## Primitive Data Types

- `String`: String: a UTF8 encoded String
- `Double`: a double precision 64 bit numeric value
- `Integer`: a 32 bit signed whole number
- `Long`: a 64 bit signed whole number
- `DateTime`: an ISO-8601 compatible time instance, with optional time zone and UTZ offset
- `Boolean`: a Boolean value, either true or false

`Xxx[]` specifies an array of `Xxx`.

## Resource Definitions

- Assets
- Participants
- Transactions
- Events
- Enumerated Types
- Concepts

A property is either a field or a relationship.

### Predefined Resource Definitions

> - *Italic*: The resource definition is abstract
> - **Bold**: The field is the identity of the resource
> - \*xxx: The field "xxx" is a relationship (pointer)
> - xxx?: the field "xxx" is optional

#### Assets (`asset`)

- *Asset*: ()
  - *Registry*: (**registryId**, name, type, system)
    - AssetRegistry
    - ParticipantRegistry
    - TransactionRegistry
  - Network: (**networkId**)
  - HistorianRecord: (**transactionId**, transactionType, \*transactionInvoked, \*participantInvoking?, \*identityUsed?, eventsEmitted?, transactionTimestamp)
  - Identity: (**identityId**, name, issuer, certificate, state, \*participant)

#### Participants (`participant`)

- *Participant*: ()
  - NetworkAdmin: (**participantId**)

#### Transactions (`transaction`)

- *Transaction*: (**transactionId**, timestamp)
  - *RegistryTransaction*: (\*targetRegistry)
    - *AssetTransaction*: (resources)
      - AddAsset: ()
      - UpdateAsset: ()
      - RemoveAsset: (resourceIds)
    - *ParticipantTransaction*: (resources)
      - AddParticipant: ()
      - UpdateParticipant: ()
      - RemoveParticipant: (resourceIds)
  - IssueIdentity: (\*participant, identityName)
  - BindIdentity: (\*participant, certificate)
  - ActivateCurrentIdentity: ()
  - RevokeIdentity: (\*identity)
  - StartBusinessNetwork: (businessNetworkArchive, logLevel?, bootstrapTransactions?)
  - ResetBusinessNetwork: ()
  - UpdateBusinessNetwork: (businessNetworkArchive)
  - SetLogLevel: (newLogLevel)

#### Events (`event`)

- Event: (**eventId**, timestamp)

#### Enumerated Types (`enum`)

- IdentityState: {ISSUED, BOUND, ACTIVATED, REVOKED}

#### Concepts (`concept`)

(none)

## References

- https://hyperledger.github.io/composer/reference/cto_language.html
- https://github.com/hyperledger/composer/blob/master/packages/composer-common/lib/system/org.hyperledger.composer.system.cto
  - Version: 37ba09c9afd95f15a40d53fd78b48dd6482e8904
