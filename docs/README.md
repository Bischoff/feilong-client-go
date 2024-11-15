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
   * 7.5.3 - `GetGuestMinidisksInfo()`
   * 7.5.4 - `AddGuestDisks()`
   * 7.5.5 - `ConfigureGuestDisks()`
   * 7.5.6 - `DeleteGuestDisks()`
   * 7.5.7 - `AttachGuestVolume()`
   * 7.5.8 - `DetachGuestVolume()`
   * 7.5.9 - `GetGuestsStats()`
   * 7.5.10 - `GetGuestsInterfaceStats()`
   * 7.5.11 - `GetGuestsNICInfo()`
   * 7.5.12 - `ShowGuestDefinition()`
   * 7.5.13 - `DeleteGuest()`
   * 7.5.14 - `GetGuestPowerStateFromHypervisor()`
   * 7.5.15 - `GetGuestInfo()`
   * 7.5.16 - `GetGuestUserDirectory()`
   * 7.5.17 - `GetGuestAdaptersInfo()`
   * 7.5.18 - `CreateGuestNIC()`
   * 7.5.19 - `CreateGuestNetworkInterface()`
   * 7.5.20 - `DeleteGuestNetworkInterface()`
   * 7.5.21 - `StartGuest()`
   * 7.5.22 - `StopGuest()`
   * 7.5.23 - `SoftStopGuest()`
   * 7.5.24 - `PauseGuest()`
   * 7.5.25 - `UnpauseGuest()`
   * 7.5.26 - `RebootGuest()`
   * 7.5.27 - `ResetGuest()`
   * 7.5.28 - `GetGuestConsoleOutput()`
   * 7.5.29 - `LiveMigrateGuest()`
   * 7.5.30 - `RegisterGuest()`
   * 7.5.31 - `DeregisterGuest()`
   * 7.5.32 - `LiveResizeGuestCPUs()`
   * 7.5.33 - `ResizeGuestCPUs()`
   * 7.5.34 - `LiveResizeGuestMemory()`
   * 7.5.35 - `ResizeGuestMemory()`
   * 7.5.36 - `DeployGuest()`
   * 7.5.37 - `CaptureGuest()`
   * 7.5.38 - `GrowGuestRootVolume()`
   * 7.5.39 - `GetGuestPowerState()`
   * 7.5.40 - `UpdateGuestNIC()`
   * 7.5.41 - `DeleteGuestNIC()`
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
 * 7.9 - [Volumes](https://github.com/Bischoff/feilong-client-go/blob/main/volumes.go)
   * 7.9.1 - `RefreshVolumeBootmapInfo()`
   * 7.9.2 - `GetVolumeConnector()`
   * 7.9.3 - `GetVolumeFCPUsage()`
   * 7.9.4 - `SetVolumeFCPUsage()`
 * 7.10 - [Files](https://github.com/Bischoff/feilong-client-go/blob/main/file.go)
   * 7.10.1 - `ImportFile()`
   * 7.10.2 - `ExportFile()`
 * 7.11 - [Switch](https://github.com/Bischoff/feilong-client-go/blob/main/switch.go)
   * 7.11.1 - `GetSwitchInfo()`
