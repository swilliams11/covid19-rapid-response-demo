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

## Agent Questions
Following the documentation below to deploy your agent and then you ask the following questions.  The agent is
is able to handle variations in the questions below, so these are sample questions.  Only a subset of bank names were entered into the agent.

### Query a real bank
#### 1. What are the hours of operation for `BANK_NAME`?

The agent will send the request to the Google Cloud Function which will convert your address to a Geocode using the **Geocoding API** and then search for the closest
location near your address with the **Google Places API**.  Enter an actual address for best results.  

**Agent Follow-up Prompt**
`What is your address?`

**User Response**
`500 W Madison St Chicago, IL`

**Response**
`The closest BANK_NAME is ADDRESS and the hours of operation are LOCATION_HOURS_OF_OPERATION`

#### 2. `BANK_NAME` open branches?

This is will always return a default response using fulfillment.
**Agent Follow-up Prompt**
`What is your address?`

**User Response**
`Dallas, TX`

**Response**
`I found some locations that are near by your address.
1234 Main Street Dallas, TX
2234 Park Place Dallas, TX
3678 Smith Way Dallas, TX`

#### 3. Can I refinance my mortgage?

The agent will respond with the default text response.  This question does not use fulfillment.  

**Response**
```Generally, you should consider refinancing if interest rates are falling or home prices are rising.  
If you are interested in refinancing, then please start the online application process here.
```

#### 4. I will miss my credit card payment, how can you help?

The agent will respond with the default text response.  This question does not use fulfillment.  

**Response**
```If you want to formally request that your credit card payment be delayed for yourself or your business, 
then please enroll online for the fastest service.
```


### Acme Bank Questions
You can query a fictious bank that will fulfill the request with a Google Cloud Function.

#### 1. What are the hours of operation for Acme bank?

This is will always return a default response using fulfillment.  
**Agent Follow-up Prompt**
`What is your address?`

**User Response**
`Chicago, IL`

**Response**
```The closest Acme bank is 333 Jackson blvd, Chicago IL and the hours of operation are 
Monday - Friday 9AM - 4PM, Saturday 9AM - 12PM and closed on Sunday
```

#### 2. `BANK_NAME` open branches?

The agent will send the request to the Google Cloud Function, which will convert your address to a Geocode using the **Geocoding API** and then search for the 3 closest
operational locations near your address with the **Google Places API**.  Enter an actual address for best results.  This will return the 
vicinity, not the actual address.  

**Agent Follow-up Prompt**
`What is your address?`

**User Response**
`Dallas, TX`

**Response**
```I found some locations that are near by your address.
ADDRESS_1
ADDRESS_2
ADDRESS_3
```

## Agent Deployment
Google provides the [COVID-19 Banking Dialogflow virtual agent 
template (the "Template")](./covid-19-agent-template.zip), so you 
can import it into your own Dialogflow agent and make changes to fit your needs.

### Import the COVID-19 Banking Dialogflow Virtual Agent Template into Your Agent
1. Download the [COVID-19 Banking Dialogflow Agent Template](./covid-19-agent-template.zip)
2. Create a new agent.
3. Click the Settings icon.
4. Select the Export and Import tab, then click the IMPORT FROM ZIP button to 
import the agent template.
![Import Agent Screenshot](../resources/import-export.png)

### Import and Deploy Fulfillment into Your Agent
*Please note there are two special intents (locations.open_branches and locations.hours_of_operations) which require you to set up the Google Places API and Geocoding API. Please follow step 8-9 to set up the APIs.*
1. Download Fulfillment from [Dialogflow COVID-19 Fulfillment.](./dialogflow-fulfillment.zip).

2. Click "Fulfillment" in the left sidebar.

3. Toggle the switch to enable the Inline Editor.
![Inline Editor Screenshot](../resources/inline-editor.png).

4. Follow the instruction on the screen to enable fulfillment via Cloud 
Functions and enable billing for the Cloud project.

5. Go to the Google Cloud Console and select Cloud Functions on the left panel.
![Cloud Function Screenshot](../resources/cloud-function.png)

6. Select the fulfillment and click Edit button.
![Cloud Function Edit Screenshot](../resources/cloud-function-edit.png)

7. Under the "Source code" section, select "ZIP upload" and upload the 
fulfillment zip file downloaded at step 1. Then select a Stage bucket (you may need to create one if it hasn't been created yet).

8. [Optional] Follow [Quickstart](https://developers.google.com/maps/gmp-get-started#quickstart) to enable 
[Google Maps Places API](https://developers.google.com/places/web-service/intro) if you haven't done so. 
Go to GCP API & Services->Credentials component to create an API key for calling the Google Maps Places API 
(More detailed instructions are listed at [Get an API Key](https://developers.google.com/places/web-service/get-api-key?hl=en_US).
![Create API Key Screenshot](../resources/create-api-key.png)

9.  Set GOOGLE_MAPS_API_KEY environment variable to the API key when deploy Cloud Function. (More details can be found at  [Cloud Function Updating Environment Variable](https://cloud.google.com/functions/docs/env-var#updating_environment_variables))
![Set Maps API Key Screenshot](../resources/set-maps-api-key.png)

## Integrate with COVID-19 Banking Virtual Agent Template

### Interact with the Dialogflow Console
Type your text query input in the Dialogflow Simulator. *Please note that custom payload of response may not show up on Dialogflow Console, you can click DIAGNOSTIC INFO to get more information about the response*.
![Dialogflow Console Screenshot](../resources/dialogflow-console.png)

### Integrate with [Dialogflow Messenger](https://cloud.google.com/dialogflow/docs/integrations/dialogflow-messenger)
Follow the [instructions here](https://github.com/GoogleCloudPlatform/covid19-rapid-response-demo#integrate-with-dialogflow-messenger)

### Integrate with this Chat App and other third-party service providers
This chat application provides a front end chat interface to a Dialogflow Agent. 
*Please note this chat app not an official Google product.* 

Please following the [instructions here](https://github.com/GoogleCloudPlatform/covid19-rapid-response-demo#integrate-with-this-chat-app)

### Release Notes


## Google Cloud Disclaimer: 
The contents of this GitHub directory are provided under the Apache 2.0 license. 
