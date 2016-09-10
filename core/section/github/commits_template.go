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

package github

const commitsTemplate = `
<div class="section-github-render">
{{if .HasAuthorStats}}
	<h3>Contributors</h3>

	<p>
		There
		{{if eq 1 .NumContributors}}is{{else}}are{{end}}
		{{.NumContributors}}
		{{if eq 1 .NumContributors}}contributor{{else}}contributors{{end}}
		across {{.RepoCount}}
		{{if eq 1 .RepoCount}} repository. {{else}} repositories. {{end}}
	</p>

	<table class="contributor-table" style="width:100%;">
		<tbody class="github">
		{{range $stats := .AuthorStats}}
			<tr>
				<td class="width-5">
					<img class="github-avatar" alt="@{{$stats.Author}}" src="{{$stats.Avatar}}" height="36" width="36">
				</td>

				<td class="width-95">
					<h6>{{$stats.Author}}</h6>
					{{if gt $stats.OpenIssues 0}}
						has been assigned {{$stats.OpenIssues}}
						{{if eq 1 $stats.OpenIssues}} issue,
					 	{{else}} issues, {{end}}
					 {{end}}
					 {{if gt $stats.ClosedIssues 0}}
						{{$stats.ClosedIssues}} have been closed,
					{{end}}

					{{if gt $stats.CommitCount 0}}
						has made {{$stats.CommitCount}}
						{{if eq 1 $stats.CommitCount}} commit {{else}} commits {{end}}
						on {{len $stats.Repos}} {{if eq 1 (len $stats.Repos)}} branch. {{else}} branches. {{end}}
						<br>
						{{range $repo := $stats.Repos}}	{{$repo}}, {{end}}
					{{end}}
				</td>

			</tr>
		{{end}}
		</tbody>
	</table>
{{end}}

{{if .HasCommits}}
	<h3>Commits</h3>
	<p> There are {{len .BranchCommits}} commits by {{.NumContributors}} contributors
		across {{.RepoCount}}
		{{if eq 1 .RepoCount}} repository. {{else}} repositories. {{end}}
	</p>
	<table class="contributor-table" style="width:100%;">
		<tbody class="github">
		{{range $commit := .BranchCommits}}
			<tr>
				<td style="width:5%;">
					<img class="github-avatar" alt="@{{$commit.Name}}" src="{{$commit.Avatar}}" height="36" width="36">
				</td>
				<td style="width:45%;">
					{{if $commit.ShowUser}}
						<h6>{{$commit.Name}}</h6>
					{{end}}
					<a class="link" href="{{$commit.URL}}">{{$commit.Message}}</a><br>
					<span class="date-meta">{{if $commit.ShowDate}}{{$commit.Date}}{{end}}</span>
				</td>
				<td style="width:55%;">
					{{if $commit.ShowBranch}}{{$commit.Repo}}:<span class="branch">{{$commit.Branch}}</span>{{end}}
					<br>
				</td>
			</tr>
		{{end}}
		</tbody>
	</table>
{{end}}

</div>
`