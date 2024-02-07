<img align="right" alt="logo" src="https://rocket-rp.fun/Libs/Img/logo.png" width="200" height="200" />

# Arma Security Extension


[![License: CC BY-ND 4.0](https://img.shields.io/badge/License-CC_BY--ND_4.0-lightgrey.svg)](https://creativecommons.org/licenses/by-nd/4.0/)
[![Go Report Card](https://goreportcard.com/badge/github.com/fairytale5571/secExt)](https://goreportcard.com/report/github.com/fairytale5571/secExt)

Developed for [Rocket Life - RIP](http://rocket-rp.fun/)  


## Disclaimer: 

**Published for informational and educational purposes.**

**The legality of using this extension is questionable.**

**Use, which I do not recommend, at your own risk.**

**Responsibility for use in production lies with you.**

### User story:

There was a need to detect multi-accounts of players and clarify the characteristics of PC components.

It was both a means to combat cheaters and other dishonest players, and also a means of collecting statistics on the target audience, both in terms of optimizing game functions and monetization.

## Слава Україні 🇺🇦

### Requirements:
Go 1.17 CGO enabled GCC Compiler  

### Donate
- BTC: ```bc1qjnvkr3asrwvh55eycucnrchftwdl82w3vsk3tt```
- ETH: ```0x4c197743E4340C5695a687882663816F4F3948Bf```
- USDT ERC20: ```0x4c197743E4340C5695a687882663816F4F3948Bf```
- USDT TRC20: ```TVpRMLNMwFoM5L2R1YN4DpeNJPkrrVeJzY```

### Build:
```bash   
 make build  
```  

### Commands:

- info
- goarch
- close
- version
- isAdmin
- get_HWID (deprecated)
- get_Process
- get_MAC
- get_IP
- get_GeoIP
- get_Sd
- v (deprecated)
- GetCPU_id
- GetCPU_name
- GetMother_id
- GetMother_name
- GetBios_id
- GetBios_ReleaseDate
- GetBios_Version
- GetRam_serialNumber
- GetRam_capacity
- GetRam_partNumber
- GetRam_Name
- GetProduct_Date
- GetProduct_Version
- Get_Drives
- get_Product
- GetPC_name
- Get_SID
- Get_VRAM_name
- get_CSP
- uuid  - generate uniquie UUID string
- setEnv - 2 arguments `[key, value]`
- getEnv - 1 arguments `[key]`

### Screenshots
- 1_c  - set credentials for GDrive
- 2_c  - set token for GDrive
- 3_c  - Make screenshot and upload, with timeout 5 second (game will be freezed on 5 sec)
- 3_c_t - Same as 3_c but without freeze
- 4_c - clear cache images

### Registry reader

Allowed categories `classes_root, current_user, local_machine, users, current_config`
Known issues: Not all items in categories can be read, in windows we have limits, use better `current_user`

- 1_r - write registry key. 4 arguments `[category, path, key, value]`
- 2_r - read registry key. 3 arguments `[category, path, key]`
- 3_r - delete registry key.  3 arguments `[category, path, key]`

### File reader

- 1_f - write file. 2 arguments `[path, content]`
- 2_f - read file. 1 argument `[path]`
- 3_f - delete file. 1 argument `[path]`

