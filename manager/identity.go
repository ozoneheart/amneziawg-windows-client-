/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019-2022 WireGuard LLC. All Rights Reserved.
 */

package manager

import (
    "errors"
    "strings"

    "golang.org/x/sys/windows/registry"

    "github.com/amnezia-vpn/amneziawg-windows/v3/conf"
)

const (
    ProductName = "MyAmneziaWG"

    ManagerServiceName  = "MyAmneziaWGManager"
    TunnelServicePrefix = "MyAmneziaWGTunnel$"

    tunnelNamePrefix = "MyAmneziaWG-"
    adminRegKey      = `Software\MyAmneziaWG`
)

func internalTunnelName(name string) string {
    if strings.HasPrefix(name, tunnelNamePrefix) {
        return name
    }
    return tunnelNamePrefix + name
}

func externalTunnelName(name string) string {
    return strings.TrimPrefix(name, tunnelNamePrefix)
}

func serviceNameOfTunnel(name string) (string, error) {
    name = internalTunnelName(name)
    if !conf.TunnelNameIsValid(name) {
        return "", errors.New("Tunnel name is not valid")
    }
    return TunnelServicePrefix + name, nil
}

func adminBool(name string) bool {
    key, err := registry.OpenKey(registry.LOCAL_MACHINE, adminRegKey, registry.QUERY_VALUE|registry.WOW64_64KEY)
    if err != nil {
        return false
    }
    defer key.Close()
    value, _, err := key.GetIntegerValue(name)
    return err == nil && value != 0
}
