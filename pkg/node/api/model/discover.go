// Copyright (C) 2014-2018 Goodrain Co., Ltd.
// RAINBOND, Application Management Platform

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package model

const (
	//PREFIX PREFIX
	PREFIX string = "PREFIX"
	//HEADERS HEADERS
	HEADERS string = "HEADERS"
	//DOMAINS DOMAINS
	DOMAINS string = "DOMAINS"
	//MaxConnections The maximum number of connections that Envoy will make to the upstream cluster. If not specified, the default is 1024.
	MaxConnections string = "MaxConnections"
	//MaxPendingRequests The maximum number of pending requests that Envoy will allow to the upstream cluster. If not specified, the default is 1024
	MaxPendingRequests string = "MaxPendingRequests"
	//MaxRequests The maximum number of parallel requests that Envoy will make to the upstream cluster. If not specified, the default is 1024.
	MaxRequests string = "MaxRequests"
	//MaxActiveRetries  The maximum number of parallel retries that Envoy will allow to the upstream cluster. If not specified, the default is 3.
	MaxActiveRetries string = "MaxActiveRetries"
	//UPSTREAM upStream
	UPSTREAM string = "upStream"
	//DOWNSTREAM downStream
	DOWNSTREAM string = "downStream"
	//WEIGHT WEIGHT
	WEIGHT string = "WEIGHT"
	//MODELWEIGHT MODEL_WEIGHT
	MODELWEIGHT string = "weight_model"
	//MODELPREFIX MODEL_PREFIX
	MODELPREFIX string = "prefix_model"
)
