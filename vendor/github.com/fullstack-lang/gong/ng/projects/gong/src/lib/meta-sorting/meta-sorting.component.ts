// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { MetaDB } from '../meta-db'
import { MetaService } from '../meta.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-meta-sorting',
  templateUrl: './meta-sorting.component.html',
  styleUrls: ['./meta-sorting.component.css']
})
export class MetaSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of Meta instances that are in the association
  associatedMetas = new Array<MetaDB>();

  constructor(
    private metaService: MetaService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of meta instances
    public dialogRef: MatDialogRef<MetaSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getMetas()
  }

  getMetas(): void {
    this.frontRepoService.pull(this.dialogData.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let meta of this.frontRepo.Metas_array) {
          let ID = this.dialogData.ID
          let revPointerID = meta[this.dialogData.ReversePointer as keyof MetaDB] as unknown as NullInt64
          let revPointerID_Index = meta[this.dialogData.ReversePointer + "_Index" as keyof MetaDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedMetas.push(meta)
          }
        }

        // sort associated meta according to order
        this.associatedMetas.sort((t1, t2) => {
          let t1_revPointerID_Index = t1[this.dialogData.ReversePointer + "_Index" as keyof typeof t1] as unknown as NullInt64
          let t2_revPointerID_Index = t2[this.dialogData.ReversePointer + "_Index" as keyof typeof t2] as unknown as NullInt64
          if (t1_revPointerID_Index && t2_revPointerID_Index) {
            if (t1_revPointerID_Index.Int64 > t2_revPointerID_Index.Int64) {
              return 1;
            }
            if (t1_revPointerID_Index.Int64 < t2_revPointerID_Index.Int64) {
              return -1;
            }
          }
          return 0;
        });
      }
    )
  }

  drop(event: CdkDragDrop<string[]>) {
    moveItemInArray(this.associatedMetas, event.previousIndex, event.currentIndex);

    // set the order of Meta instances
    let index = 0

    for (let meta of this.associatedMetas) {
      let revPointerID_Index = meta[this.dialogData.ReversePointer + "_Index" as keyof MetaDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedMetas.forEach(
      meta => {
        this.metaService.updateMeta(meta, this.dialogData.GONG__StackPath)
          .subscribe(meta => {
            this.metaService.MetaServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer + ' done');
  }
}