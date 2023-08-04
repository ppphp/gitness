// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package space

import (
	"net/http"

	"github.com/harness/gitness/internal/api/controller/space"
	"github.com/harness/gitness/internal/api/render"
	"github.com/harness/gitness/internal/api/request"
	"github.com/harness/gitness/types"
	"github.com/harness/gitness/types/enum"
)

// HandleListRepos writes json-encoded list of repos in the request body.
func HandleListPipelines(spaceCtrl *space.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		session, _ := request.AuthSessionFrom(ctx)
		spaceRef, err := request.GetSpaceRefFromPath(r)
		if err != nil {
			render.TranslatedUserError(w, err)
			return
		}

		// TODO: Use pipeline filter
		filter := request.ParseRepoFilter(r)
		if filter.Order == enum.OrderDefault {
			filter.Order = enum.OrderAsc
		}

		repos, totalCount, err := spaceCtrl.ListPipelines(ctx, session, spaceRef, &types.PipelineFilter{})
		if err != nil {
			render.TranslatedUserError(w, err)
			return
		}

		render.Pagination(r, w, filter.Page, filter.Size, int(totalCount))
		render.JSON(w, http.StatusOK, repos)
	}
}
