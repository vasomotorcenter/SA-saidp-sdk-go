package oath

import (
	"testing"

	factors "github.com/jhickmansa/saidp-sdk-go/services/factors"
	sa "github.com/secureauthcorp/saidp-sdk-go"
)

/*
**********************************************************************
*   @author jhickman@secureauth.com
*
*  Copyright (c) 2017, SecureAuth
*  All rights reserved.
*
*    Redistribution and use in source and binary forms, with or without modification,
*    are permitted provided that the following conditions are met:
*
*    1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
*
*    2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer
*    in the documentation and/or other materials provided with the distribution.
*
*    3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived
*    from this software without specific prior written permission.
*
*    THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO,
*    THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR
*    CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
*    PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
*    LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE,
*    EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
**********************************************************************
 */

const (
	fAppID  = ""
	fAppKey = ""
	fHost   = ""
	fRealm  = ""
	fPort   = 443
	fUser   = ""
	fPass   = ""
)

var (
	fOtp = "62150185"
	fID  = ""
)

func TestOathSettings(t *testing.T) {
	client, err := sa.NewClient(fAppID, fAppKey, fHost, fPort, fRealm, true, false)
	if err != nil {
		t.Error(err)
	}
	factorsRequest := new(factors.Request)
	factorsResponse, err := factorsRequest.Get(client, fUser)
	if err != nil {
		t.Error(err)
	}
	valid, err := factorsResponse.IsSignatureValid(client)
	if err != nil {
		t.Error(err)
	}
	if !valid {
		t.Error("Response signature is invalid")
	}

	for _, factor := range factorsResponse.Factors {
		if factor.FactorType == "oath" {
			fID = factor.ID
		}
	}

	oathRequest := new(Request)
	oathResponse, err := oathRequest.GetOATHSettings(client, fUser, fPass, fOtp, fID)
	if err != nil {
		t.Error(err)
	}
	valid, err = oathResponse.IsSignatureValid(client)
	if err != nil {
		t.Error(err)
	}
	if !valid {
		t.Error("Response signature is invalid")
	}
}
