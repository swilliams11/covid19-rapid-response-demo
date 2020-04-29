// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Component, OnInit } from '@angular/core';
import { ChatService, Message } from '../chat.service';
import { AgentService, Agent } from '../agent.service';
import { Observable } from 'rxjs/Observable';
import { Title } from '@angular/platform-browser';
import 'rxjs/add/operator/scan';


@Component({
  selector: 'chat-dialog',
  templateUrl: './chat-dialog.component.html',
  styleUrls: ['./chat-dialog.component.scss']
})

export class ChatDialogComponent implements OnInit {

  messages: Observable<Message[]>;
  agent: Agent = { title: "" };

  constructor(public chat: ChatService, public agentService: AgentService, private titleService: Title) { }

  ngOnInit() {
    // appends to array after each new message is added to feedSource
    this.messages = this.chat.conversation.asObservable().scan((acc, val) => acc.concat(val));
    this.chat.converseText("default hello");
    this.chat.conversation.subscribe(value => {
      this.hideTemp();
    })
    this.getTitle();

  }

  getTitle(): void {
    this.agentService.get().subscribe(agent => {
      this.agent = agent;
      this.setTitle(agent.title);
    });
  }

  setTitle(newTitle: string): void {
    this.titleService.setTitle(newTitle);
  }



  ngAfterViewChecked() {
  }



  hideTemp() {
    var temps = document.querySelectorAll(".message.temp");

    temps.forEach(function (temp: HTMLElement) {
      temp.style.display = "none";
    });
  }

}
