// Copyright 2016 Documize Inc. <legal@documize.com>. All rights reserved.
//
// This software (Documize Community Edition) is licensed under
// GNU AGPL v3 http://www.gnu.org/licenses/agpl-3.0.en.html
//
// You can operate outside the AGPL restrictions by purchasing
// Documize Enterprise Edition and obtaining a commercial license
// by contacting <sales@documize.com>.
//
// https://documize.com

import { hash } from 'rsvp';
import { inject as service } from '@ember/service';
import Route from '@ember/routing/route';
import AuthenticatedRouteMixin from 'ember-simple-auth/mixins/authenticated-route-mixin';

export default Route.extend(AuthenticatedRouteMixin, {
	documentService: service('document'),
	linkService: service('link'),
	folderService: service('folder'),
	userService: service('user'),

	model() {
		let document = this.modelFor('document').document;
		this.browser.setTitle(document.get('name'));
		this.browser.setMetaDescription(document.get('excerpt'));

		return hash({
			folders: this.modelFor('document').folders,
			folder: this.modelFor('document').folder,
			document: this.modelFor('document').document,
			pages: this.get('documentService').getPages(this.modelFor('document').document.get('id')),
			links: this.modelFor('document').links,
			sections: this.modelFor('document').sections,
			permissions: this.modelFor('document').permissions
		});
	},

	setupController(controller, model) {
		controller.set('folders', model.folders);
		controller.set('folder', model.folder);
		controller.set('document', model.document);
		controller.set('pages', model.pages);
		controller.set('links', model.links);
		controller.set('sections', model.sections);
		controller.set('permissions', model.permissions);
	}
});