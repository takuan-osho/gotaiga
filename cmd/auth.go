// Copyright © 2018 SHIMIZU Taku <shimizu.taku@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewAuthCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Manage auth API",
		Long: `Manage auth API

	References:
	* https://taigaio.github.io/taiga-doc/dist/api.html#_auth
	* https://taigaio.github.io/taiga-doc/dist/api.html#auth
	`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(NewAuthLoginCmd())

	return cmd
}
func NewAuthLoginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "login a user",
		Long: `login a user

		API References:
		* https://taigaio.github.io/taiga-doc/dist/api.html#auth
		`,
		RunE: runAuthLoginCmd,
	}
	return cmd
}

func runAuthLoginCmd(cmd *cobra.Command, args []string) error {
	client, err := NewDefaultClient()
	if err != nil {
		return errors.Wrap(err, "NewDefaultClient failed")
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeoutTime)
	defer cancel()

	userAuthDetail, err := client.AuthLogin(ctx)
	if err != nil {
		return errors.Wrap(err, "AuthLogin failed")
	}

	cmd.Println(userAuthDetail.AuthToken)
	return nil
}

func (client *Client) AuthLogin(ctx context.Context) (*UserAuthDetail, error) {
	subPath := "auth"
	payload := &NormalLogin{
		Type:     "normal",
		Username: client.Username,
		Password: client.Password,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	body := strings.NewReader(bytes.NewBuffer(b).String())

	req, err := client.NewRequest(ctx, "POST", subPath, body)
	if err != nil {
		return nil, err
	}

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("Status Code Error: status code is %v", resp.StatusCode)
	}

	var userAuthDetailResponse UserAuthDetail
	if err := decodeBody(resp, &userAuthDetailResponse, nil); err != nil {
		return nil, err
	}

	return &userAuthDetailResponse, nil
}
