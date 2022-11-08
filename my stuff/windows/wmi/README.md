# Play with WMI / CIM

> get parameters for the query (in powershell)

```powershell
# everything
Get-WmiObject -list
Get-CimClass
# or a subset of cmdlets
Get-WmiObject -list | findstr.exe /I route
```

> Get parameters for the app

```powershell
# use structure in the app
Get-WmiObject Win32_IP4RouteTable
(get-cimclass –class Win32_IP4RouteTable).CimClassProperties
```
