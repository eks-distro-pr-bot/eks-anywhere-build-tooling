## **Kubernetes vSphere Cloud Provider**
![1.27 Version](https://img.shields.io/badge/1--27%20version-v1.27.0-blue)
![1.28 Version](https://img.shields.io/badge/1--28%20version-v1.28.1-blue)
![1.29 Version](https://img.shields.io/badge/1--29%20version-v1.29.2-blue)
![1.30 Version](https://img.shields.io/badge/1--30%20version-v1.30.2-blue)
![1.31 Version](https://img.shields.io/badge/1--31%20version-v1.31.1-blue)
![Build Status](https://codebuild.us-west-2.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoiYzQ3dzRvZHVqU2MvYnVuMzB3QmRZdVd1U1RabVorWnlqTXBYUGxDSGk2NXJXUU12c3pLQ25CQUdaQmlNUE84S0JIVVZUU0ozeTJJb3J0NWxNejNSbzk4PSIsIml2UGFyYW1ldGVyU3BlYyI6IkhLNTZwQ0hiZDZVUzVRdXYiLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=main)

[Cloud Provider vSphere](https://github.com/kubernetes/cloud-provider-vsphere) defines the vSphere-specific implementation of the Kubernetes controller-manager. The Cloud Provider Interface (CPI) allows customers to run Kubernetes clusters on vSphere infrastructure. It replaces the Kubernetes Controller Manager for only the cloud-specific control loops. The CPI integration connects to vCenter Server and maps information about the infrastructure, such as VMs, disks, and so on, back to the Kubernetes API.

You can find the latest version of this image [on ECR Public Gallery](https://gallery.ecr.aws/eks-anywhere/kubernetes/cloud-provider-vsphere/cpi/manager).
