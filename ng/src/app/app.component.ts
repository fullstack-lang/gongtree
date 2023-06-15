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
  tree = "Tree"
  view = this.tree

  views: string[] = [this.tree, this.default];

  GONG__StackPath = "tree"

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
