BASEDIR = $(shell pwd)
PROJECTAPPENGINE=$(COVID_CHAT_APPENGINE_PROJECT)
PROJECTDIALOGFLOW=$(COVID_CHAT_DIALOGFLOW_PROJECT)

PROJECTNUMBER=$(shell gcloud projects list --filter="$(PROJECTAPPENGINE)" --format="value(PROJECT_NUMBER)")


env:
	gcloud config set project $(PROJECTAPPENGINE)

check:
	@echo ~~~~~~~~~~~~~ Testing project ids.
	@echo COVID_CHAT_APPENGINE_PROJECT=$(COVID_CHAT_APPENGINE_PROJECT)
	@echo COVID_CHAT_DIALOGFLOW_PROJECT=$(COVID_CHAT_DIALOGFLOW_PROJECT)


clean:
	-rm -rf server/dist		

frontend: clean
	cd chat-interface && ng build --prod


deploy: env clean frontend
	cd server && gcloud app deploy -q

deploy-stage: env clean
	cd chat-interface && ng build --configuration stage
	cd server && gcloud app deploy -q --version stage --no-promote	

build: env clean 
	gcloud builds submit --config cloudbuild.yaml .	

init: env apis appengine serviceaccount secret permissions
	@echo ~~~~~~~~~~~~~ Install node_modules. 
	-cd chat-interface && npm install
	@echo ~~~~~~~~~~~~~Install Go vendor dependencies
	-cd server && go mod vendor
	@echo ~~~~~~~~~~~~~ Create Angular builder for Cloud Build 
	-cd builder && make build

apis:
	@echo ~~~~~~~~~~~~~ Enable API access on $(PROJECTAPPENGINE)
	-gcloud services enable cloudbuild.googleapis.com
	-gcloud services enable appengine.googleapis.com 
	-gcloud services enable dialogflow.googleapis.com
	-gcloud services enable compute.googleapis.com
	-gcloud services enable secretmanager.googleapis.com

appengine: 	
	@echo ~~~~~~~~~~~~~ Intialize AppEngine on $(PROJECTAPPENGINE)
	-gcloud app create --region us-central

permissions:
	@echo ~~~~~~~~~~~~~ Enable Cloud Run service account to deploy to AppEngine on $(PROJECTAPPENGINE)
	-gcloud projects add-iam-policy-binding $(PROJECTAPPENGINE) \
  	--member serviceAccount:$(PROJECTNUMBER)@cloudbuild.gserviceaccount.com \
  	--role roles/appengine.appAdmin
	@echo ~~~~~~~~~~~~~ Enable AppEngine on $(PROJECTAPPENGINE) service account to call Dialogflow on $(PROJECTDIALOGFLOW)	  
	-gcloud projects add-iam-policy-binding $(PROJECTDIALOGFLOW) \
  	--member serviceAccount:$(PROJECTAPPENGINE)@appspot.gserviceaccount.com \
  	--role roles/dialogflow.client
	@echo ~~~~~~~~~~~~~ Grant service account access to $(PROJECTDIALOGFLOW)
	-gcloud projects add-iam-policy-binding $(PROJECTDIALOGFLOW) \
  	--member serviceAccount:chat-sa@$(PROJECTDIALOGFLOW).iam.gserviceaccount.com \
  	--role roles/dialogflow.client 
	@echo ~~~~~~~~~~~~~ Enable AppEngine on $(PROJECTAPPENGINE) service account access to secrets
	-gcloud secrets add-iam-policy-binding PROJECTDIALOGFLOW \
    --member serviceAccount:$(PROJECTAPPENGINE)@appspot.gserviceaccount.com \
    --role roles/secretmanager.secretAccessor
	@echo ~~~~~~~~~~~~~ Grant service account access to secrets
	-gcloud secrets add-iam-policy-binding PROJECTDIALOGFLOW \
    --member serviceAccount:chat-sa@$(PROJECTDIALOGFLOW).iam.gserviceaccount.com \
    --role roles/secretmanager.secretAccessor      


serviceaccount:
	@echo ~~~~~~~~~~~~~ Create service account for Dialogflow   
	-gcloud iam service-accounts create chat-sa \
    --description "A service account for development of frontend of a Dialogflow agent" \
    --display-name "Dialogflow Chat Bot" --project $(PROJECTDIALOGFLOW)
	@echo ~~~~~~~~~~~~~ Grant service account access to $(PROJECTDIALOGFLOW)
	-gcloud projects add-iam-policy-binding $(PROJECTDIALOGFLOW) \
  	--member serviceAccount:chat-sa@$(PROJECTDIALOGFLOW).iam.gserviceaccount.com \
  	--role roles/dialogflow.admin
	@echo ~~~~~~~~~~~~~ Grant service account access to secrets
	-gcloud secrets add-iam-policy-binding PROJECTDIALOGFLOW \
    --member serviceAccount:chat-sa@$(PROJECTDIALOGFLOW).iam.gserviceaccount.com \
    --role roles/secretmanager.secretAccessor       
	@echo ~~~~~~~~~~~~~ Download key for service account. 
	-gcloud iam service-accounts keys create creds/creds.json \
  	--iam-account chat-sa@$(PROJECTDIALOGFLOW).iam.gserviceaccount.com  

cleansa:
	-gcloud iam service-accounts delete \
	chat-sa@$(PROJECTDIALOGFLOW).iam.gserviceaccount.com \
	--project $(PROJECTDIALOGFLOW) -q

secret:
	@echo ~~~~~~~~~~~~~ Creating secret to store reference to $(PROJECTDIALOGFLOW) 
	-gcloud secrets create "PROJECTDIALOGFLOW" --replication-policy="automatic"
	-echo $(PROJECTDIALOGFLOW) | gcloud secrets versions add "PROJECTDIALOGFLOW" --data-file=-	

dev:
	(trap 'kill 0' SIGINT; \
	cd server && \
	export PROJECTDIALOGFLOW=$(PROJECTDIALOGFLOW) && \
	export GOOGLE_APPLICATION_CREDENTIALS=$(BASEDIR)/creds/creds.json && \
	go run main.go & \
	cd $(BASEDIR)/chat-interface && ng serve --open )

service:
	
	cd server && \
	export PROJECTDIALOGFLOW=$(PROJECTDIALOGFLOW) && \
	export GOOGLE_APPLICATION_CREDENTIALS=$(BASEDIR)/creds/creds.json && \
	go run main.go 
	
	
