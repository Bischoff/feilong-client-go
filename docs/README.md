# Documentation - Go Client Library for Feilong


## Naming Conventions

The following conventions are used both for function names and for input and output structures:

 * names use camel case, with no underscores: `user_profile` in Feilong JSON API becomes `UserProfile` in this Go library
 * acronyms are completly capitalized: `api_version` becomes `APIVersion`
 * abbreviations are sometimes completed, but never fully capitalized: `max_version` becomes `MaxVersion`, and `modID` becomes `ModuleId`
 * the Feilong name is not completely respected if it does not respect English grammar or is otherwise wrong: `GetGuestsList` becomes `GetGuestList`.

The names differ quite significantly from the ones used in the Python library.

Please refer to the individual definitions to know the exact names used in this library.


## Implemented Functions

The numbers below refer to the section numbers in the [Feilong API documentation](https://cloudlib4zvm.readthedocs.io/en/latest/restapi.html).

 * 7.2 - [Version](https://github.com/Bischoff/feilong-client-go/blob/main/version.go)
   * 7.2.1 - `GetVersion()`
 * 7.3 - [Token](https://github.com/Bischoff/feilong-client-go/blob/main/token.go)
   * 7.3.1 - `CreateToken()`
 * 7.4 - [SMAPI](https://github.com/Bischoff/feilong-client-go/blob/main/smapi.go)
   * 7.4.1 - `SMAPIHealth()`
 * 7.5 - [Guests](https://github.com/Bischoff/feilong-client-go/blob/main/guests.go)
   * 7.5.1 - `ListGuests()`
   * 7.5.2 - `CreateGuest()`
   * 7.5.3 - `GuestAddDisks()`
   * 7.5.4 - `GuestConfigureDisks()`
   * 7.5.5 - `GuestDeleteDisks()`
   * 7.5.15 - `ShowGuestDefinition()`
   * 7.5.16 - `DeleteGuest()`
   * 7.5.18 - `GetGuestInfo()`
   * 7.5.20 - `GetGuestAdaptersInfo()`
   * 7.5.21 - `CreateGuestNIC()`
   * 7.5.24 - `StartGuest()`
   * 7.5.25 - `StopGuest()`
   * 7.5.39 - `DeployGuest()`
   * 7.5.43 - `UpdateGuestNIC()`
 * 7.6 - [Host](https://github.com/Bischoff/feilong-client-go/blob/main/host.go)
   * 7.6.1 - `GetGuestList()`
   * 7.6.2 - `GetHostInfo()`
   * 7.6.3 - `GetHostDiskPoolInfo()`, `GetHostDiskPoolDetails()`
   * 7.6.4 - `GetHostDiskPoolVolumeNames()`
   * 7.6.5 - `GetHostVolumeInfo()`
   * 7.6.6 - `GetHostSSIClusterInfo()`
 * 7.7 - [Images](https://github.com/Bischoff/feilong-client-go/blob/main/images.go)
   * 7.7.1 - `ListImages()`
   * 7.7.2 - `CreateImage()`
   * 7.7.3 - `ExportImage()`
   * 7.7.4 - `GetRootDiskSize()`
   * 7.7.5 - `DeleteImage()`
 * 7.8 - [VSwitches](https://github.com/Bischoff/feilong-client-go/blob/main/vswitches.go)
   * 7.8.1 - `CreateVSwitch()`
   * 7.8.2 - `ListVSwitches()`
   * 7.8.3 - `GetVSwitchDetails()`
   * 7.8.4 - `GrantUserToVSwitch()`
   * 7.8.5 - `RevokeUserFromVSwitch()`
   * 7.8.6 - `SetUserVLANIdToVSwitch()`
   * 7.8.7 - `DeleteVSwitch()`
 * 7.9 - [Files](https://github.com/Bischoff/feilong-client-go/blob/main/file.go)
   * 7.9.1 - `ImportFile()`
   * 7.9.2 - `ExportFile()`
