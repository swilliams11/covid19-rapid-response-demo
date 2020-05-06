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

### General Questions
#### 1. Can I refinance my mortgage?

The agent will respond with the default text response.  This question does not use fulfillment.  

##### Agent Response
```
Generally, you should consider refinancing if interest rates are falling or home prices are rising.  
If you are interested in refinancing, then please start the online application process here.
```

#### 2. I will miss my credit card payment, how can you help?

The agent will respond with the default text response.  This question does not use fulfillment.  

##### Agent Response
```
If you want to formally request that your credit card payment be delayed for yourself or your business, 
then please enroll online for the fastest service.
```

#### 3. How are you assisting banking customers affected by the COVID19 pandemic?

The agent will respond with the default text response; no fulfillment.

##### Agent Response
```
We are here to help in any way we can.  First visit the CDC for the most updated information on COVID-19.  
We offer a variety of mobile services such as banking online, transferring money via our mobile application, 
depositing checks from your home by phone or our app, and the ability to make purchases with our digital wallet.  
These service should help you maintain social distancing.  We are also limiting our branch office hours and allowing 
drive through service, when available.
```

#### 4. What is the best way to bank if I am sheltering in place?

The agent will respond with the default text response; no fulfillment.

##### Agent Response
```
We offer expanded services through the use of drive-up teller windows, 
or provide assistance at ATMs located outside of the branch office. 
Please visit our website and register for online banking or download our mobile app.  
You can complete transactions such as depositing a check to your bank account or paying bills. 
In addition, you may want to consider signing up for direct deposit so that a paycheck or public 
benefits payment goes directly into your account.
```

#### 5. Are you waiving late charges?

The agent will respond with the default text response; no fulfillment.

##### Agent Response
```
We understand during this time that you may be seriously affected by COVID-19 and we may 
be able to assist you with your payment. We can provide some options depending on if it is 
a loan or credit card.  If its a loan, then we may allow you to skip loan payments with no 
adverse consequences for you, extend loan terms, or restructure the loan.  
If it is a credit card then we may be able to waive the late charge and possibly the 
interest for a set time period.  Please use our website to request more information.
```

#### 6. Can I defer loan payments because of COVID-19?

The agent will respond with the default text response; no fulfillment.

##### Agent Response
```
If you are unable to make a loan payment due to COVID-19, then we may be able to assist you. 
Please contact us for more information or visit our website for more information.
```

#### 7. How are we assisting mortgage customers affected by the COVID-19 pandemic?

The agent will respond with the default text response; no fulfillment.

##### Agent Response
```
A new federal law, the Coronavirus Aid, Relief, and Economic Security (CARES) Act, puts in place two 
protections for homeowners with federally backed mortgages:

First, your lender or loan servicer may not foreclose on you for 60 days after March 18, 2020. Specifically, 
the CARES Act prohibits lenders and servicers from beginning a judicial or non-judicial foreclosure against you, 
or from finalizing a foreclosure judgment or sale, during this period of time.

Second, if you experience financial hardship due to the coronavirus pandemic, you have a right to 
request a forbearance for up to 180 days. You also have the right to request an extension for up to another 180 days. 
You must contact your loan servicer to request this forbearance. There will be no additional fees, penalties or 
additional interest (beyond scheduled amounts) added to your account. You do not need to submit additional documentation 
to qualify other than your claim to have a pandemic-related financial hardship.  

If your loan is not federally backed then there may be other options available for your state and servicer.
```

### Query an actual banks hours of operation or branch locations
#### 1. What are the hours of operation for `BANK_NAME`?

The agent will send the request to the Google Cloud Function which will convert your address to a Geocode using the **Geocoding API** and then search for the closest
location near your address with the **Google Places API**.  Enter an actual address for best results.  

##### Agent Follow-up Prompt

`What is your address?`

##### User Response

`500 W Madison St Chicago, IL`

##### Agent Response

`The closest BANK_NAME is ADDRESS and the hours of operation are LOCATION_HOURS_OF_OPERATION`

#### 2. `BANK_NAME` open branches?

The agent will send the request to the Google Cloud Function, which will convert your address to a Geocode using the **Geocoding API** and then search for the 3 closest
operational locations near your address with the **Google Places API**.  Enter an actual address for best results.  This will return the 
vicinity, not the actual address.  

##### Agent Follow-up Prompt

`What is your address?`

##### User Response

`1600 Amphitheatre Pwky, Mountain View CA`

##### Agent Response

```
I found some locations that are near by your address.
ADDRESS_1
ADDRESS_2
ADDRESS_3
```


### Acme Bank Questions
You can query a fictious bank that will fulfill the request with a Google Cloud Function.

#### 1. What are the hours of operation for Acme bank?

This is will always return a default response using fulfillment.  

##### Agent Follow-up Prompt

`What is your address?`

##### User Response

`Chicago, IL`

##### Agent Response

```
The closest Acme bank is 333 Jackson blvd, Chicago IL and the hours of operation are 
Monday - Friday 9AM - 4PM, Saturday 9AM - 12PM and closed on Sunday
```

#### 2. `BANK_NAME` open branches?

This is will always return a default response using fulfillment.

##### Agent Follow-up Prompt

`What is your address?`

##### User Response

`Dallas, TX`

##### Agent Response

```
I found some locations that are near by your address.
1234 Main Street Dallas, TX
2234 Park Place Dallas, TX
3678 Smith Way Dallas, TX
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
![Import Agent Screenshot](../../resources/import-export.png)

### Import and Deploy Fulfillment into Your Agent
*Please note there are two special intents (locations.open_branches and locations.hours_of_operations) which require you to set up the Google Places API and Geocoding API. Please follow step 8-9 to set up the APIs.*
1. Download Fulfillment from [Dialogflow COVID-19 Fulfillment.](./dialogflow-fulfillment.zip).

2. Click "Fulfillment" in the left sidebar.

3. Toggle the switch to enable the Inline Editor.
![Inline Editor Screenshot](../../resources/inline-editor.png).

4. Follow the instruction on the screen to enable fulfillment via Cloud 
Functions and enable billing for the Cloud project.

5. Go to the Google Cloud Console and select Cloud Functions on the left panel.
![Cloud Function Screenshot](../../resources/cloud-function.png)

6. Select the fulfillment and click Edit button.
![Cloud Function Edit Screenshot](../../resources/cloud-function-edit.png)

7. Under the "Source code" section, select "ZIP upload" and upload the 
fulfillment zip file downloaded at step 1. Then select a Stage bucket (you may need to create one if it hasn't been created yet).

8. [Optional] Follow [Quickstart](https://developers.google.com/maps/gmp-get-started#quickstart) to enable 
[Google Maps Places API](https://developers.google.com/places/web-service/intro) if you haven't done so. 
Go to GCP API & Services->Credentials component to create an API key for calling the Google Maps Places API 
(More detailed instructions are listed at [Get an API Key](https://developers.google.com/places/web-service/get-api-key?hl=en_US).
![Create API Key Screenshot](../../resources/create-api-key.png)

9.  Set GOOGLE_MAPS_API_KEY environment variable to the API key when deploy Cloud Function. (More details can be found at  [Cloud Function Updating Environment Variable](https://cloud.google.com/functions/docs/env-var#updating_environment_variables))
![Set Maps API Key Screenshot](../../resources/set-maps-api-key.png)

## Integrate with COVID-19 Banking Virtual Agent Template

### Interact with the Dialogflow Console
Type your text query input in the Dialogflow Simulator. *Please note that custom payload of response may not show up on Dialogflow Console, you can click DIAGNOSTIC INFO to get more information about the response*.
![Dialogflow Console Screenshot](../../resources/dialogflow-console.png)

### Integrate with [Dialogflow Messenger](https://cloud.google.com/dialogflow/docs/integrations/dialogflow-messenger)
Follow the [instructions here](https://github.com/GoogleCloudPlatform/covid19-rapid-response-demo#integrate-with-dialogflow-messenger)

### Integrate with this Chat App and other third-party service providers
This chat application provides a front end chat interface to a Dialogflow Agent. 
*Please note this chat app not an official Google product.* 

Please following the [instructions here](https://github.com/GoogleCloudPlatform/covid19-rapid-response-demo#integrate-with-this-chat-app)

### Release Notes


## Google Cloud Disclaimer: 
The contents of this GitHub directory are provided under the Apache 2.0 license. 
