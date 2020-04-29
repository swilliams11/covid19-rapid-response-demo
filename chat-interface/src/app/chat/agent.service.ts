import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs/Observable';

export interface Agent{
  title:string;
  

}


@Injectable({
  providedIn: 'root'
})
export class AgentService {

  constructor(private http: HttpClient) { }

  get():Observable<Agent> {
      let apiURL = environment.apihost + "/agent";
      return this.http.get<Agent>(apiURL)

  }
}
