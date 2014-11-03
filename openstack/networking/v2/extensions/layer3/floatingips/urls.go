package floatingips

import "github.com/rackspace/gophercloud"

const resourcePath = "floatingips"
const novaNetResourcePath = "os-floating-ips"

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}
/*/v2/{tenant_id}/os-floating-ips*/
func resourceNovaNetFloatingIpURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(novaNetResourcePath)
}

/*
   {
    "addFloatingIp": {
        "address": "10.10.10.1"
    }}
*/
func addNovaNetFloatingIpURL(client *gophercloud.ServiceClient, server_id string) string {
	return client.ServiceURL("servers", server_id, "action")
}
