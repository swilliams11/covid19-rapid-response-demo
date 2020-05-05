<!--
  Licensed to the Apache Software Foundation (ASF) under one or more
  contributor license agreements.  See the NOTICE file distributed with
  this work for additional information regarding copyright ownership.
  The ASF licenses this file to You under the Apache License, Version 2.0
  (the "License"); you may not use this file except in compliance with
  the License.  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
-->
# COVID-19 Banking Virtual Agent Template 
Google has launched a sample Dialogflow 
template for banking institutions enabling chat or voice bots to 
answer questions about branch locations and office hours. This should help banking institutions
understand how to integrate Dialogflow into their backend systems.

It can:
* Answer questions about open branches
* Return the hours of operation for the closest branch
* Return generic responses regarding late payments and refinancing your mortgage

You can:
* Launch your bot faster with curated content
* Customize to match your website and add content to address your organization 
and community needs  

## Agent Deployment
Google provides the [COVID-19 Banking Dialogflow virtual agent 
template (the "Template")](./agent-template-banking/covid-19-agent-template.zip), so you 
can import it into your own Dialogflow agent and make changes to fit your needs.

### Import the COVID-19 Banking Dialogflow Virtual Agent Template into Your Agent
1. Download the [COVID-19 Banking Dialogflow Agent Template](./covid-19-agent-template.zip)
1. Create a new agent.
1. Click the Settings icon.
1. Select the Export and Import tab, then click the IMPORT FROM ZIP button to 
import the agent template.
![Import Agent Screenshot](./resources/import-export.png)

### Import and Deploy Fulfillment into Your Agent
*Please note there are two special intents (locations.open_branches and locations.hours_of_operations) which require you to set up the Google Places API and Geocoding API. Please follow step 8-9 to set up the APIs.*
1. Download Fulfillment from [Dialogflow COVID-19 Fulfillment.](./dialogflow-fulfillment.zip).
2. Click "Fulfillment" in the left sidebar.
3. Toggle the switch to enable the Inline Editor.
![Inline Editor Screenshot](./resources/inline-editor.png).
4. Follow the instruction on the screen to enable fulfillment via Cloud 
Functions and enable billing for the Cloud project.
5. Go to the Google Cloud Console and select Cloud Functions on the left panel.
![Cloud Function Screenshot](./resources/cloud-function.png)
6. Select the fulfillment and click Edit button.
![Cloud Function Edit Screenshot](./resources/cloud-function-edit.png)
7. Under the "Source code" section, select "ZIP upload" and upload the 
fulfillment zip file downloaded at step 1. Then select a Stage bucket (you may need to create one if it hasn't been created yet).
8. [Optional] Follow [Quickstart](https://developers.google.com/maps/gmp-get-started#quickstart) to enable 
[Google Maps Places API](https://developers.google.com/places/web-service/intro) if you haven't done so. 
Go to GCP API & Services->Credentials component to create an API key for calling the Google Maps Places API 
(More detailed instructions are listed at [Get an API Key](https://developers.google.com/places/web-service/get-api-key?hl=en_US).
![Create API Key Screenshot](./resources/create-api-key.png)
9.  Set GOOGLE_MAPS_API_KEY environment variable to the API key when deploy Cloud Function. (More details can be found at  [Cloud Function Updating Environment Variable](https://cloud.google.com/functions/docs/env-var#updating_environment_variables))
![Set Maps API Key Screenshot](./resources/set-maps-api-key.png)

## Integrate with COVID-19 Banking Virtual Agent Template

### Interact with the Dialogflow Console
Type your text query input in the Dialogflow Simulator. *Please note that custom payload of response may not show up on Dialogflow Console, you can click DIAGNOSTIC INFO to get more information about the response*.
![Dialogflow Console Screenshot](./resources/dialogflow-console.png)

### Integrate with [Dialogflow Messenger](https://cloud.google.com/dialogflow/docs/integrations/dialogflow-messenger)
1. Go to the Dialogflow Console.
1. Select your agent.
1. Click Integrations in the left sidebar menu.
1. Enable the Dialogflow Messenger integration.
1. Open the agent's web page using the provided link. This page provides a text 
chat interface. Type your input query and press enter. The agent responds with 
the response from your agent.
![Dialogflow Messenger Screenshot](./resources/dialogflow-messenger.png)
1. [Optional] You may also embed the Dialogflow Messenger into your website by 
following the instructions.

### Integrate with this Chat App and other third-party service providers
This chat application provides a front end chat interface to a Dialogflow Agent. 
*Please note this chat app not an official Google product.* 

Please following the [instructions here](https://github.com/GoogleCloudPlatform/covid19-rapid-response-demo#integrate-with-this-chat-app)

### Release Notes


## Google Cloud Disclaimer: 
The contents of this GitHub directory are provided under the Apache 2.0 license. 
