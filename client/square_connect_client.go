// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/client/apple_pay"
	"github.com/jefflinse/square-connect/client/bank_accounts"
	"github.com/jefflinse/square-connect/client/cash_drawers"
	"github.com/jefflinse/square-connect/client/catalog"
	"github.com/jefflinse/square-connect/client/checkout"
	"github.com/jefflinse/square-connect/client/customer_groups"
	"github.com/jefflinse/square-connect/client/customer_segments"
	"github.com/jefflinse/square-connect/client/customers"
	"github.com/jefflinse/square-connect/client/devices"
	"github.com/jefflinse/square-connect/client/disputes"
	"github.com/jefflinse/square-connect/client/employees"
	"github.com/jefflinse/square-connect/client/inventory"
	"github.com/jefflinse/square-connect/client/invoices"
	"github.com/jefflinse/square-connect/client/labor"
	"github.com/jefflinse/square-connect/client/locations"
	"github.com/jefflinse/square-connect/client/loyalty"
	"github.com/jefflinse/square-connect/client/merchants"
	"github.com/jefflinse/square-connect/client/mobile_authorization"
	"github.com/jefflinse/square-connect/client/o_auth"
	"github.com/jefflinse/square-connect/client/orders"
	"github.com/jefflinse/square-connect/client/payments"
	"github.com/jefflinse/square-connect/client/refunds"
	"github.com/jefflinse/square-connect/client/subscriptions"
	"github.com/jefflinse/square-connect/client/team"
	"github.com/jefflinse/square-connect/client/terminal"
	"github.com/jefflinse/square-connect/client/transactions"
	"github.com/jefflinse/square-connect/client/v1_employees"
	"github.com/jefflinse/square-connect/client/v1_items"
	"github.com/jefflinse/square-connect/client/v1_locations"
	"github.com/jefflinse/square-connect/client/v1_transactions"
)

// Default square connect HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "connect.squareup.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new square connect HTTP client.
func NewHTTPClient(formats strfmt.Registry) *SquareConnect {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new square connect HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *SquareConnect {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new square connect client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *SquareConnect {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(SquareConnect)
	cli.Transport = transport
	cli.ApplePay = apple_pay.New(transport, formats)
	cli.BankAccounts = bank_accounts.New(transport, formats)
	cli.CashDrawers = cash_drawers.New(transport, formats)
	cli.Catalog = catalog.New(transport, formats)
	cli.Checkout = checkout.New(transport, formats)
	cli.CustomerGroups = customer_groups.New(transport, formats)
	cli.CustomerSegments = customer_segments.New(transport, formats)
	cli.Customers = customers.New(transport, formats)
	cli.Devices = devices.New(transport, formats)
	cli.Disputes = disputes.New(transport, formats)
	cli.Employees = employees.New(transport, formats)
	cli.Inventory = inventory.New(transport, formats)
	cli.Invoices = invoices.New(transport, formats)
	cli.Labor = labor.New(transport, formats)
	cli.Locations = locations.New(transport, formats)
	cli.Loyalty = loyalty.New(transport, formats)
	cli.Merchants = merchants.New(transport, formats)
	cli.MobileAuthorization = mobile_authorization.New(transport, formats)
	cli.OAuth = o_auth.New(transport, formats)
	cli.Orders = orders.New(transport, formats)
	cli.Payments = payments.New(transport, formats)
	cli.Refunds = refunds.New(transport, formats)
	cli.Subscriptions = subscriptions.New(transport, formats)
	cli.Team = team.New(transport, formats)
	cli.Terminal = terminal.New(transport, formats)
	cli.Transactions = transactions.New(transport, formats)
	cli.V1Employees = v1_employees.New(transport, formats)
	cli.V1Items = v1_items.New(transport, formats)
	cli.V1Locations = v1_locations.New(transport, formats)
	cli.V1Transactions = v1_transactions.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// SquareConnect is a client for square connect
type SquareConnect struct {
	ApplePay apple_pay.ClientService

	BankAccounts bank_accounts.ClientService

	CashDrawers cash_drawers.ClientService

	Catalog catalog.ClientService

	Checkout checkout.ClientService

	CustomerGroups customer_groups.ClientService

	CustomerSegments customer_segments.ClientService

	Customers customers.ClientService

	Devices devices.ClientService

	Disputes disputes.ClientService

	Employees employees.ClientService

	Inventory inventory.ClientService

	Invoices invoices.ClientService

	Labor labor.ClientService

	Locations locations.ClientService

	Loyalty loyalty.ClientService

	Merchants merchants.ClientService

	MobileAuthorization mobile_authorization.ClientService

	OAuth o_auth.ClientService

	Orders orders.ClientService

	Payments payments.ClientService

	Refunds refunds.ClientService

	Subscriptions subscriptions.ClientService

	Team team.ClientService

	Terminal terminal.ClientService

	Transactions transactions.ClientService

	V1Employees v1_employees.ClientService

	V1Items v1_items.ClientService

	V1Locations v1_locations.ClientService

	V1Transactions v1_transactions.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *SquareConnect) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.ApplePay.SetTransport(transport)
	c.BankAccounts.SetTransport(transport)
	c.CashDrawers.SetTransport(transport)
	c.Catalog.SetTransport(transport)
	c.Checkout.SetTransport(transport)
	c.CustomerGroups.SetTransport(transport)
	c.CustomerSegments.SetTransport(transport)
	c.Customers.SetTransport(transport)
	c.Devices.SetTransport(transport)
	c.Disputes.SetTransport(transport)
	c.Employees.SetTransport(transport)
	c.Inventory.SetTransport(transport)
	c.Invoices.SetTransport(transport)
	c.Labor.SetTransport(transport)
	c.Locations.SetTransport(transport)
	c.Loyalty.SetTransport(transport)
	c.Merchants.SetTransport(transport)
	c.MobileAuthorization.SetTransport(transport)
	c.OAuth.SetTransport(transport)
	c.Orders.SetTransport(transport)
	c.Payments.SetTransport(transport)
	c.Refunds.SetTransport(transport)
	c.Subscriptions.SetTransport(transport)
	c.Team.SetTransport(transport)
	c.Terminal.SetTransport(transport)
	c.Transactions.SetTransport(transport)
	c.V1Employees.SetTransport(transport)
	c.V1Items.SetTransport(transport)
	c.V1Locations.SetTransport(transport)
	c.V1Transactions.SetTransport(transport)
}
