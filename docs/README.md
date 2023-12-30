# Documentation - Go Client Library for Feilong


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
   * 7.5.3 - `AddGuestDisks()`
   * 7.5.4 - `ConfigureGuestDisks()`
   * 7.5.5 - `DeleteGuestDisks()`
   * 7.5.6 - TODO
   * 7.5.7 - TODO
   * 7.5.8 - TODO
   * 7.5.9 - TODO
   * 7.5.10 - TODO
   * 7.5.11 - TODO
   * 7.5.12 - `GetGuestsStats()`
   * 7.5.13 - `GetGuestsInterfaceStats()`
   * 7.5.14 - `GetGuestsNICInfo()`
   * 7.5.15 - `ShowGuestDefinition()`
   * 7.5.16 - `DeleteGuest()`
   * 7.5.17 - `GetGuestPowerStateFromHypervisor()`
   * 7.5.18 - `GetGuestInfo()`
   * 7.5.19 - `GetGuestUserDirectory()`
   * 7.5.20 - `GetGuestAdaptersInfo()`
   * 7.5.21 - `CreateGuestNIC()`
   * 7.5.22 - `CreateGuestNetworkInterface()`
   * 7.5.23 - `DeleteGuestNetworkInterface()`
   * 7.5.24 - `StartGuest()`
   * 7.5.25 - `StopGuest()`
   * 7.5.26 - `SoftStopGuest()`
   * 7.5.27 - `PauseGuest()`
   * 7.5.28 - `UnpauseGuest()`
   * 7.5.29 - `RebootGuest()`
   * 7.5.30 - `ResetGuest()`
   * 7.5.31 - `GetGuestConsoleOutput()`
   * 7.5.32 - `LiveMigrateGuest()`
   * 7.5.33 - `RegisterGuest()`
   * 7.5.34 - `DeregisterGuest()`
   * 7.5.35 - `LiveResizeGuestCPUs()`
   * 7.5.36 - `ResizeGuestCPUs()`
   * 7.5.37 - `LiveResizeGuestMemory()`
   * 7.5.38 - `ResizeGuestMemory()`
   * 7.5.39 - `DeployGuest()`
   * 7.5.40 - `CaptureGuest()`
   * 7.5.41 - `GrowGuestRootVolume()`
   * 7.5.42 - `GetGuestPowerState()`
   * 7.5.43 - `UpdateGuestNIC()`
   * 7.5.44 - `DeleteGuestNIC()`
 * 7.6 - [Host](https://github.com/Bischoff/feilong-client-go/blob/main/host.go)
   * 7.6.1 - `GetHostGuestList()`
   * 7.6.2 - `GetHostInfo()`
   * 7.6.3 - `GetHostDiskPoolInfo()`, `GetHostDiskPoolDetails()`
   * 7.6.4 - `GetHostDiskPoolVolumeNames()`
   * 7.6.5 - `GetHostVolumeInfo()`
   * 7.6.6 - `GetHostSSIClusterInfo()`
 * 7.7 - [Images](https://github.com/Bischoff/feilong-client-go/blob/main/images.go)
   * 7.7.1 - `ListImages()`
   * 7.7.2 - `CreateImage()`
   * 7.7.3 - `ExportImage()`
   * 7.7.4 - `GetImageRootDiskSize()`
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
