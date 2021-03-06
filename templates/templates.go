// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

import (
	"io"
	"log"
	"html/template"
)

var tmpls = make(map[string]*template.Template)

func init() {
	tmpls["index.html"] = template.Must(template.New("base").Parse(baseSrc))
	template.Must(tmpls["index.html"].New("index").Parse(indexSrc))

	tmpls["login.html"] = template.Must(template.New("base").Parse(baseSrc))
	template.Must(tmpls["login.html"].New("login").Parse(loginSrc))

	tmpls["signup.html"] = template.Must(template.New("base").Parse(baseSrc))
	template.Must(tmpls["signup.html"].New("signup").Parse(signupSrc))

	tmpls["forgotpass.html"] = template.Must(template.New("base").Parse(baseSrc))
	template.Must(tmpls["forgotpass.html"].New("forgotpass").Parse(forgotpassSrc))

	tmpls["resetpass.html"] = template.Must(template.New("base").Parse(baseSrc))
	template.Must(tmpls["resetpass.html"].New("resetpass").Parse(resetpassSrc))

	tmpls["profile.html"] = template.Must(template.New("base").Parse(baseSrc))
	template.Must(tmpls["profile.html"].New("profile").Parse(profileSrc))

	tmpls["admin.html"] = template.Must(template.New("base").Parse(baseSrc))
	template.Must(tmpls["admin.html"].New("admin").Parse(adminSrc))

	tmpls["adminusers.html"] = template.Must(template.New("base").Parse(baseSrc))
	template.Must(tmpls["adminusers.html"].New("adminusers").Parse(adminUsersSrc))

	tmpls["admingroups.html"] = template.Must(template.New("base").Parse(baseSrc))
	template.Must(tmpls["admingroups.html"].New("admingroups").Parse(adminGroupsSrc))

	tmpls["admingroupmembers.html"] = template.Must(template.New("base").Parse(baseSrc))
	template.Must(tmpls["admingroupmembers.html"].New("admingroupmembers").Parse(adminGroupMembersSrc))

	tmpls["accessdenied.html"] = template.Must(template.New("base").Parse(baseSrc))
	template.Must(tmpls["accessdenied.html"].New("accessdenied").Parse(accessDeniedSrc))

	tmpls["pagelist.html"] = template.Must(template.New("base").Parse(baseSrc))
	template.Must(tmpls["pagelist.html"].New("pagelist").Parse(pageListSrc))

	tmpls["filelist.html"] = template.Must(template.New("base").Parse(baseSrc))
	template.Must(tmpls["filelist.html"].New("filelist").Parse(fileListSrc))

	tmpls["search.html"] = template.Must(template.New("base").Parse(baseSrc))
	template.Must(tmpls["search.html"].New("search").Parse(searchSrc))
}

func Render(wr io.Writer, template string, data interface{}) {
	err := tmpls[template].Execute(wr, data)
	if err != nil {
		log.Panicf("[ERROR] Error rendering %s: %s\n", template, err)
	}
}