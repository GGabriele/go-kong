# Table of Contents

- [v0.22.0](#v0220---20210922)
- [v0.21.0](#v0210---20210826)
- [v0.20.0](#v0200---20210707)
- [v0.19.0](#v0190---20210514)
- [v0.18.0](#v0180---20210505)
- [v0.17.0](#v0170---20210405)
- [v0.16.0](#v0160---20210303)
- [v0.15.0](#v0150---20210122)
- [v0.14.0](#v0140---20210112)
- [v0.13.0](#v0130---20200804)
- [v0.12.0](#v0120---20200730)
- [v0.11.0](#v0110---20200117)
- [v0.10.0](#v0100---20191027)
- [v0.9.0](#v090---20190824)
- [v0.8.0](#v080---20190821)
- [v0.7.0](#v070---20190813)
- [v0.6.2](#v062---20190809)
- [v0.6.1](#v061---20190809)
- [v0.6.0](#v060---20190809)
- [v0.5.1](#v051---20190805)
- [v0.5.0](#v050---20190607)
- [v0.4.1](#v041---20190411)
- [v0.4.0](#v040---20190406)
- [0.3.0](#030---20181219)
- [0.2.0](#020---20181219)
- [0.1.0](#010---20181201)

## [v0.22.0] - 2021/09/22

### Added

- Test that target updates work as expected.
  [#85](https://github.com/Kong/go-kong/pull/85)
- Workspace service supports new `ExistsByName()` API. This is similar to
  `Exists()`, but it only accepts names, not IDs. This function can check if a
  workspace exists even if the RBAC user has access to that namespace only.
  `Exists()` requires access to the default namespace.
  [#90](https://github.com/Kong/go-kong/pull/90)

### Fixed

- The generic `exists()` client functions uses GETs instead of HEADs. Kong
  <=2.6 respond to HEAD requests with 200s as long as the endpoint has a valid
  form. They do not check if entities actually exist, and `exists()` would
  return true for entities that don't exist because of this.
  [#90](https://github.com/Kong/go-kong/pull/90)

## [v0.21.0] - 2021/08/26

### Added

- oauth2 entities support `hash_secret`.
  [#74](https://github.com/Kong/go-kong/pull/74)

### Fixed

- Plugin validation checks against the correct status code.
  [#81](https://github.com/Kong/go-kong/pull/81)

## [v0.20.0] - 2021/07/07

### Added

- FriendlyName is now defined for entities based on deck types.
  [#68](https://github.com/Kong/go-kong/pull/68)
- Added Info service for interacting with information exposed by the admin API root endpoint.
  [#65](https://github.com/Kong/go-kong/pull/65)
- Implemented client-level support for Kong workspaces.
  [#62](https://github.com/Kong/go-kong/pull/62)

### Changed

- Internally wrapped errors now use the standard Go library instead of a 3rd party
  wrapping lib and consequently several string versions of errors have changed.
  [#66](https://github.com/Kong/go-kong/pull/66)
- Various improvements to repository CI: caching go modules, updating codegen, e.t.c.

## [v0.19.0] - 2021/05/05

### Added

- Client now allows reading the version of Kong. [#48](https://github.com/Kong/go-kong/pull/48)

## [v0.18.0] - 2021/05/05

### Added

- Workspace and tag existence functions. [#56](https://github.com/Kong/go-kong/pull/56)
- Plugin schema retrieval. [#57](https://github.com/Kong/go-kong/pull/57)

### Changed

- The default client now sets a default timeout. [#51](https://github.com/Kong/go-kong/pull/51)
- Optimized HTTP header handling. [#49](https://github.com/Kong/go-kong/pull/49)
- Changed return type of `kong.HTTPClientWithHeaders(...)`: now returns `*http.Client` instead of `http.Client`. [#46](https://github.com/Kong/go-kong/pull/46)
- Now uses Go 1.16 module semantics. [#40](https://github.com/Kong/go-kong/pull/40)
- Now uses Dependabot to manage GitHub Actions. [#41](https://github.com/Kong/go-kong/pull/41), [#42](https://github.com/Kong/go-kong/pull/42)
- Now uses code-generator v0.21.0. [#43](https://github.com/Kong/go-kong/pull/43)
- CI no longer uses Bintray repositories. [#55](https://github.com/Kong/go-kong/pull/55)
- Dropped testing against older versions of Kong. [#58](https://github.com/Kong/go-kong/pull/58)

## [v0.17.0] - 2021/04/05

### Added

- Added support for Developer entities. Thanks to @ChristianJacquot!
  [#27](https://github.com/Kong/go-kong/pull/27)
- Added support for Developer Role entities. Thanks to @mmorel-35!
  [#30](https://github.com/Kong/go-kong/pull/27)
- RBAC roles now support pagination and listing all entities. Thanks to
  @mmorel-35!
  [#30](https://github.com/Kong/go-kong/pull/27)
- Tests can now require the Portal. Added helpers to enable the Portal and
  related settings.
  [#30](https://github.com/Kong/go-kong/pull/27)
- Clients now use interfaces. Other libraries that use go-kong can define mock
  types that implement these interfaces for unit tests that do not require an
  actual Kong instance.
  [#24](https://github.com/Kong/go-kong/pull/27)

### Fixed

- RBAC roles now properly include their `negative` field in requests to Kong.
  [#32](https://github.com/Kong/go-kong/pull/27)

## [v0.16.0] - 2021/03/03

### Added

- `Service` now includes `URL`.

## [v0.15.0] - 2021/01/22

### Added

- `Route` now includes `RequestBuffering` and `ResponseBuffering`.

## [v0.14.0] - 2021/01/12

### Breaking changes

- HTTP error format changed to feature HTTP codes

### Added

- RBACUser support
- Support for auto-expiring key-auth key TTL
- Support for RBAC roles
- Support for RBAC permissions
- DeepCopy annotations for enterprise entities

## [v0.13.0] - 2020/08/04

### Summary

This release renames the package from `github.com/hbagdi/go-kong` to
`github.com/kong/go-kong`.

## [v0.12.0] - 2020/07/30

### Summary

This release adds support for Kong 2.1 series and
a number of enterprise entities.

### Added

- Added `HTTPClientWithHeaders` helper function to inject authn/authz
  related HTTP headers in requests to kong.

- New fields in types:
  - `Service` struct has three new fields:
    - `TLSVerifyDepth`
    - `TLSVerify`
    - `CACertificates`
  - `ClientCertificate` field has been added to `Upstream` struct.
  - `Type` field has been added to `PassiveHealthcheck` struct.

- Following new services for Kong Enterprise have been introduced:
  - WorkspaceService to manage workspaces in kong
  - AdminService to manage users of Kong Manager
  - MTLSAuthService to manage MTLS credentials

### Misc

- Changed the branch name from `master` to `main`
- Introduced linters to improve code health

## [v0.11.0] - 2020/01/17

### Summary

- This release adds support for Kong 2.0.0.

### Added

- `Threshold` field has been added to Upstream struct.
- `PathHandling` field has been added to Route struct.

## [v0.10.0] - 2019/10/27

### Summary

- This release adds support for Kong 1.4.

### Added

- `HostHeader` field has been added to Upstream struct.
- `Tags` field has been added to the following types:
  - `KeyAuth`
  - `Basicauth`
  - `HMACAuth`
  - `Oauth2Credential`
  - `ACLGroup`
  - `JWTAuth`

## [v0.9.0] - 2019/08/24

### Breaking changes

- `client.Do()` returns a response object even on errors so that
  clients can inspect the response directly when needed.
  The error condition has changed to include HTTP status codes 300 to 399
  as success and not failure.
  [b1f9eda31](https://github.com/hbagdi/go-kong/commit/b1f9eda311e1d9c9d6b0f5a5e33a3d399d85faf6)
- `String()` method has been dropped from all types defined in this package.

### Added

- `NewRequest()` method helping with creating HTTP requests is now exported
- URLs parsed inside the package are more robust.
- New method `GetByCustomID` has been introduced to fetch Consumers by
  `custom_id`.

## [v0.8.0] - 2019/08/21

### Added

- Oauth2Credential type and service has been added
  which can be used to create Oauth2 credentials in Kong for some
  Oauth2 flows.

## [v0.7.0] - 2019/08/13

### Summary

This release brings support for CRUD methods for
authentication credentials in Kong.

### Added

- The following credentials and corresponding services have been added:
  - `key-auth`
  - `basic-auth`
  - `hmac-auth`
  - `jwt`
  - `acl`

## [v0.6.2] - 2019/08/09

### Fix

- Add missing omitempty tag to ClientCertificate field

## [v0.6.1] - 2019/08/09

### Fix

- Fix a typo in Service struct definition for YAML tag

## [v0.6.0] - 2019/08/09

### Summary

- This release adds support for Kong 1.3.

### Breaking change

- `Validator` Interface has been dropped and Valid() method from
  all entities is dropped.
  [#8](https://github.com/hbagdi/go-kong/pull/8)

### Added

- `Headers` field has been added to Route struct.
  [#5](https://github.com/hbagdi/go-kong/pull/5)
- `ClientCertificate` field has been added to Service struct.
  [#5](https://github.com/hbagdi/go-kong/pull/5)
- `CACertificate` is a new core entity in Kong. A struct to represent
  it and a corresponding new service is added.
  [#5](https://github.com/hbagdi/go-kong/pull/5)
- `Algorithm` field has been added to Upstream struct.
  [#9](https://github.com/hbagdi/go-kong/pull/9)

## [v0.5.1] - 2019/08/05

### Fix

- Add missing healthchecks.active.unhealthy.interval field to Upstream
  [#6](https://github.com/hbagdi/go-kong/issues/6)

## [v0.5.0] - 2019/06/07

### Summary

- This release adds support for Kong 1.2.

### Added

- Added HTTPSRedirectStatusCode property to Route struct.
  [#3](https://github.com/hbagdi/go-kong/pull/3)

### Breaking change

- `Create()` for Custom Entities now supports HTTP PUT method.
  If `id` is specified in the object, it will be used to PUT the entity.
  This was always POST previously.
  [#3](https://github.com/hbagdi/go-kong/pull/3)

## [v0.4.1] - 2019/04/11

### Fix

- Add `omitempty` property to Upstream fields for Kong 1.0 compatibility

## [v0.4.0] - 2019/04/06

### Summary

- This release adds support for features released in Kong 1.1.
  This version is compatible with Kong 1.0 and Kong 1.1.

### Breaking Change

- Please note that the version naming scheme for this library has changed from
  `x.y.z` to `vX.Y.Z`. This is to ensure compatibility with Go modules.

### Added

- `Tags` field has been added to all Kong Core entity structs.
- List methods now support tag based filtering introduced in Kong 1.1.
  Tags can be ANDed or ORed together. `ListOpt` struct can be used to
  specify the tags for filtering.
- `Protocols` field has been added to Plugin struct.
- New fields `Type`, `HTTPSSni` and `HTTPSVerifyCertificate` have been
  introduced for Active HTTPS healthchecks.
- `TargetService` has two new methods `MarkHealthy()` and `MarkUnhealthy()`
  to change the health of a target.

## [0.3.0] - 2018/12/19

### Summary

- This release adds support for Kong 1.0.
  It is not compatible with 0.x.y  versions of Kong due to breaking
  Admin API changes as the deprecated API entity is dropped.
- The code and API for the library is same as 0.2.0, with the exception
  that struct defs and services related to `API` is dropped.

### Breaking changes

- `API` struct definition is no longer available.
- `APIService` is no longer available. Please ensure your code doesn't rely
  on these before upgrading.
- `Plugin` struct has dropped the `API` field.

## [0.2.0] - 2018/12/19

### Summary

- This release adds support for Kong 0.15.x.
  It is not compatible with any other versions of Kong due to breaking
  Admin API changes in Kong for Plugins, Upstreams and Targets entities.

### Breaking changes

- `Target` struct now has an `Upstream` member in place of `UpstreamID`.
- `Plugin` struct now has `Consumer`, `API`, `Route`, `Service` members
  instead of `ConsumerID`, `APIID`, `RouteID` and `ServiceID`.

### Added

- `RunOn` property has been added to `Plugin`.
- New properties are added to `Route` for L4 proxy support.

## [0.1.0] - 2018/12/01

### Summary

- Debut release of this library
- This release comes with support for Kong 0.14.x
- The library is not expected to work with previous or later
  releases of Kong since every release of Kong is introducing breaking changes
  to the Admin API.

[v0.22.0]: https://github.com/Kong/go-kong/compare/v0.21.0...v0.22.0
[v0.21.0]: https://github.com/Kong/go-kong/compare/v0.20.0...v0.21.0
[v0.20.0]: https://github.com/Kong/go-kong/compare/v0.19.0...v0.20.0
[v0.19.0]: https://github.com/Kong/go-kong/compare/v0.18.0...v0.19.0
[v0.18.0]: https://github.com/Kong/go-kong/compare/v0.17.0...v0.18.0
[v0.17.0]: https://github.com/Kong/go-kong/compare/v0.16.0...v0.17.0
[v0.16.0]: https://github.com/Kong/go-kong/compare/v0.15.0...v0.16.0
[v0.15.0]: https://github.com/Kong/go-kong/compare/v0.14.0...v0.15.0
[v0.14.0]: https://github.com/Kong/go-kong/compare/v0.13.0...v0.14.0
[v0.13.0]: https://github.com/Kong/go-kong/compare/v0.12.0...v0.13.0
[v0.12.0]: https://github.com/Kong/go-kong/compare/v0.11.0...v0.12.0
[v0.11.0]: https://github.com/Kong/go-kong/compare/v0.10.0...v0.11.0
[v0.10.0]: https://github.com/Kong/go-kong/compare/v0.9.0...v0.10.0
[v0.9.0]: https://github.com/Kong/go-kong/compare/v0.8.0...v0.9.0
[v0.8.0]: https://github.com/Kong/go-kong/compare/v0.7.0...v0.8.0
[v0.7.0]: https://github.com/Kong/go-kong/compare/v0.6.2...v0.7.0
[v0.6.2]: https://github.com/Kong/go-kong/compare/v0.6.1...v0.6.2
[v0.6.1]: https://github.com/Kong/go-kong/compare/v0.6.0...v0.6.1
[v0.6.0]: https://github.com/Kong/go-kong/compare/v0.5.1...v0.6.0
[v0.5.1]: https://github.com/Kong/go-kong/compare/v0.5.0...v0.5.1
[v0.5.0]: https://github.com/Kong/go-kong/compare/v0.4.1...v0.5.0
[v0.4.1]: https://github.com/Kong/go-kong/compare/v0.4.0...v0.4.1
[v0.4.0]: https://github.com/Kong/go-kong/compare/0.3.0...v0.4.0
[0.3.0]: https://github.com/Kong/go-kong/compare/0.2.0...0.3.0
[0.2.0]: https://github.com/Kong/go-kong/compare/0.1.0...0.2.0
[0.1.0]: https://github.com/Kong/go-kong/compare/87666c7fe73477d1874d35d690301241cd23059f...0.1.0
