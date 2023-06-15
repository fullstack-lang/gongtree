import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as gongtree from 'gongtree'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  default = 'Gongtree Data/Model'
  view = this.default

  views: string[] = [this.default];

  GONG__StackPath = "github.com/fullstack-lang/gongtree/go/models"

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
