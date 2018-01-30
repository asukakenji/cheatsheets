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
> - __Underline__: The field is a relationship (pointer)
> - xxx?: the field "xxx" is optional

#### Assets (`asset`)

- *Asset*: ()
  - *Registry*: (**registryId**, name, type, system)
    - AssetRegistry
    - ParticipantRegistry
    - TransactionRegistry
  - Network: (**networkId**)
  - HistorianRecord: (**transactionId**, transactionType, __transactionInvoked__, __participantInvoking__?, __identityUsed__?, eventsEmitted?, transactionTimestamp)
  - Identity: (**identityId**, name, issuer, certificate, state, __participant__)

#### Participants (`participant`)

- *Participant*: ()
  - NetworkAdmin: (**participantId**)

#### Transactions (`transaction`)

- *Transaction*: (**transactionId**, timestamp)
  - *RegistryTransaction*: (__targetRegistry__)
    - *AssetTransaction*: (resources)
      - AddAsset: ()
      - UpdateAsset: ()
      - RemoveAsset: (resourceIds)
    - *ParticipantTransaction*: (resources)
      - AddParticipant: ()
      - UpdateParticipant: ()
      - RemoveParticipant: (resourceIds)
  - IssueIdentity: (_participant_, identityName)
  - BindIdentity: (_participant_, certificate)
  - ActivateCurrentIdentity: ()
  - RevokeIdentity: (_identity_)
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
