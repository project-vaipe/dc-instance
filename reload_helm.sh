#!/bin/bash
# Copyright 2022 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
CUSTOM_DC_DOMAIN="vaipe-datacommons.com"
CONTACT_EMAIL="nam.abc4444@gmail.com"
PROJECT_ID="vaipe-383102"




ROOT=$PWD

# Clone DC website repo and mixer submodule.
#rm -rf website
#git clone https://github.com/datacommonsorg/website --branch $CUSTOM_DC_RELEASE_TAG --single-branch

cd website
WEBSITE_GITHASH=$(git rev-parse --short=7 HEAD)

# Always force Mixer submodule to be cloned.
#rm -rf mixer
#git submodule foreach git pull origin master
#git submodule update --init --recursive

WEBSITE_ROOT=$PWD



# Install k8s resources using helm.
CLUSTER_LOCATION="asia-southeast1-a"

cd $ROOT

 mkdir -p yq
 cd yq
 wget https://github.com/mikefarah/yq/releases/download/v4.33.3/yq_linux_amd64.tar.gz -O yq.tar.gz
 tar xvf yq.tar.gz
 sudo mv yq_linux_amd64 /usr/local/bin/yq
 cd ..


sh $WEBSITE_ROOT/scripts/deploy_gke_helm.sh \
  -e vaipe \
  -l asia-southeast1-a


_success_msg="
###############################################################################
# Status: Successfully launched the installer in $PROJECT_ID.
###############################################################################


###############################################################################
# Action required:
Please don't forget to email support+custom@datacommons.org with your
GCP project id for data access.
###############################################################################


###############################################################################
# Action required:
Please also make sure to click on the activation email for $DOMAIN
If the contact email has previously been used to verify domains,
then $DOMAIN will be already active without needing activation emails.
###############################################################################


To check the status of $DOMAIN, please visit the link below.
https://console.cloud.google.com/net-services/domains/registrations/list?project=$PROJECT_ID

Note:
You should expect the instance to be accessible via $DOMAIN
within 30 minutes or so after support@datacommons.org responds.
"
echo "$_success_msg"