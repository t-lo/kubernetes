/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was automatically generated by lister-gen

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	certificates "k8s.io/kubernetes/pkg/apis/certificates"
)

// CertificateSigningRequestLister helps list CertificateSigningRequests.
type CertificateSigningRequestLister interface {
	// List lists all CertificateSigningRequests in the indexer.
	List(selector labels.Selector) (ret []*certificates.CertificateSigningRequest, err error)
	// Get retrieves the CertificateSigningRequest from the index for a given name.
	Get(name string) (*certificates.CertificateSigningRequest, error)
	CertificateSigningRequestListerExpansion
}

// certificateSigningRequestLister implements the CertificateSigningRequestLister interface.
type certificateSigningRequestLister struct {
	indexer cache.Indexer
}

// NewCertificateSigningRequestLister returns a new CertificateSigningRequestLister.
func NewCertificateSigningRequestLister(indexer cache.Indexer) CertificateSigningRequestLister {
	return &certificateSigningRequestLister{indexer: indexer}
}

// List lists all CertificateSigningRequests in the indexer.
func (s *certificateSigningRequestLister) List(selector labels.Selector) (ret []*certificates.CertificateSigningRequest, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*certificates.CertificateSigningRequest))
	})
	return ret, err
}

// Get retrieves the CertificateSigningRequest from the index for a given name.
func (s *certificateSigningRequestLister) Get(name string) (*certificates.CertificateSigningRequest, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(certificates.Resource("certificatesigningrequest"), name)
	}
	return obj.(*certificates.CertificateSigningRequest), nil
}
