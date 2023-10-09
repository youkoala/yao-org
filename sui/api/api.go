package api

import "github.com/yaoapp/gou/api"

var dsl = []byte(`
{
	"name": "SUI API",
	"description": "The API for SUI",
	"version": "1.0.0",
	"guard": "-",
	"group": "__yao/sui/v1",
	"paths": [
		{
			"path": "/:id/template",
			"method": "GET",
			"process": "sui.Template.Get",
			"in": ["$param.id"],
			"out": { "status": 200, "type": "application/json" }
		},{
			"path": "/:id/template/:template_id",
			"method": "GET",
			"process": "sui.Template.Find",
			"in": ["$param.id", "$param.template_id"],
			"out": { "status": 200, "type": "application/json" }
		},

		{
			"path": "/:id/locale/:template_id",
			"method": "GET",
			"process": "sui.Locale.Get",
			"in": ["$param.id", "$param.template_id"],
			"out": { "status": 200, "type": "application/json" }
		},{
			"path": "/:id/theme/:template_id",
			"method": "GET",
			"process": "sui.Theme.Get",
			"in": ["$param.id", "$param.template_id"],
			"out": { "status": 200, "type": "application/json" }
		},

		{
			"path": "/:id/block/:template_id",
			"method": "GET",
			"process": "sui.Block.Get",
			"in": ["$param.id", "$param.template_id"],
			"out": { "status": 200, "type": "application/json" }
		},{
			"path": "/:id/block/:template_id/:block_id",
			"method": "GET",
			"process": "sui.Block.Find",
			"in": ["$param.id", "$param.template_id", "$param.block_id"],
			"out": { "status": 200, "type": "text/javascript" }
		},

		{
			"path": "/:id/component/:template_id",
			"method": "GET",
			"process": "sui.Component.Get",
			"in": ["$param.id", "$param.template_id"],
			"out": { "status": 200, "type": "application/json" }
		},{
			"path": "/:id/component/:template_id/:component_id",
			"method": "GET",
			"process": "sui.Component.Find",
			"in": ["$param.id", "$param.template_id", "$param.component_id"],
			"out": { "status": 200, "type": "text/javascript" }
		},
		
		{
			"path": "/:id/page/:template_id/*route",
			"method": "GET",
			"process": "sui.Page.Get",
			"in": ["$param.id", "$param.template_id", "$param.route"],
			"out": { "status": 200, "type": "application/json" }
		},{
			"path": "/:id/page/tree/:template_id/*route",
			"method": "GET",
			"process": "sui.Page.Tree",
			"in": ["$param.id", "$param.template_id", "$param.route"],
			"out": { "status": 200, "type": "application/json" }
		},{
			"path": "/:id/page/save/:template_id/*route",
			"method": "POST",
			"process": "sui.Page.Save",
			"in": ["$param.id", "$param.template_id", "$param.route", ":context"],
			"out": { "status": 200, "type": "application/json" }
		},{
			"path": "/:id/page/temp/:template_id/*route",
			"method": "POST",
			"process": "sui.Page.SaveTemp",
			"in": ["$param.id", "$param.template_id", "$param.route", ":context"],
			"out": { "status": 200, "type": "application/json" }
		},{
			"path": "/:id/page/create/:template_id/*route",
			"method": "POST",
			"process": "sui.Page.Create",
			"in": ["$param.id", "$param.template_id", "$param.route", ":context"],
			"out": { "status": 200, "type": "application/json" }
		},{
			"path": "/:id/page/exist/:template_id/*route",
			"method": "GET",
			"process": "sui.Page.Exist",
			"in": ["$param.id", "$param.template_id", "$param.route"],
			"out": { "status": 200, "type": "application/json" }
		},{
			"path": "/:id/page/remove/:template_id/*route",
			"method": "POST",
			"process": "sui.Page.Remove",
			"in": ["$param.id", "$param.template_id", "$param.route"],
			"out": { "status": 200, "type": "application/json" }
		},
		
		{
			"path": "/:id/editor/render/:template_id/*route",
			"method": "GET",
			"process": "sui.Editor.Render",
			"in": ["$param.id", "$param.template_id", "$param.route", ":query"],
			"out": { "status": 200, "type": "application/json" }
		},{
			"path": "/:id/editor/render/:template_id/*route",
			"method": "POST",
			"process": "sui.Editor.RenderAfterSaveTemp",
			"in": ["$param.id", "$param.template_id", "$param.route", ":context", ":query"],
			"out": { "status": 200, "type": "application/json" }
		},{
			"path": "/:id/editor/:kind/source/:template_id/*route",
			"method": "GET",
			"process": "sui.Editor.Source",
			"in": ["$param.id", "$param.template_id", "$param.route", "$param.kind", ":query"],
			"out": { "status": 200, "type": "application/json" }
		},{
			"path": "/:id/editor/:kind/source/:template_id/*route",
			"method": "POST",
			"process": "sui.Editor.SourceAfterSaveTemp",
			"in": ["$param.id", "$param.template_id", "$param.route", ":context", "$param.kind", ":query"],
			"out": { "status": 200, "type": "application/json" }
		},

		{
			"path": "/:id/asset/:template_id/@assets/*path",
			"method": "GET",
			"process": "sui.Template.Asset",
			"in": ["$param.id", "$param.template_id", "$param.path"],
			"out": {
				"status": 200,
				"body": "?:content",
				"headers": { "Content-Type": "?:type"}
			}
		},{
			"path": "/:id/asset/:template_id/@pages/*path",
			"method": "GET",
			"process": "sui.Page.Asset",
			"in": ["$param.id", "$param.template_id", "$param.path"],
			"out": {
				"status": 200,
				"body": "?:content",
				"headers": { "Content-Type": "?:type"}
			}
		},

		{
			"path": "/:id/preview/:template_id/*route",
			"method": "GET",
			"process": "sui.Preview.Render",
			"in": ["$param.id", "$param.template_id", "$param.path", "$header.Referer", "$query.r", "$query.t"],
			"out": {"status": 200, "type": "text/html; charset=utf-8"}
		}
	],
}
`)

func registerAPI() error {
	_, err := api.LoadSource("<sui.v1>.yao", dsl, "sui.v1")
	return err
}
