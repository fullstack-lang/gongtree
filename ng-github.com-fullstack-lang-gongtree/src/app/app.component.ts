import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

import { AngularSplitModule } from 'angular-split';

import { TreeComponent } from '../../projects/gongtreespecific/src/public-api';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  standalone: true,
  imports: [
    MatRadioModule,
    FormsModule,
    CommonModule,
    MatButtonModule,
    MatIconModule,

    AngularSplitModule,

    TreeComponent,
  ]
})
export class AppComponent implements OnInit {

  default = 'Gongtree Data/Model'
  tree = "Tree"
  view = this.tree

  views: string[] = [this.tree, this.default];

  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  StackName = "gongtree"
  StackType = "github.com/fullstack-lang/gongtree/go/models"

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}