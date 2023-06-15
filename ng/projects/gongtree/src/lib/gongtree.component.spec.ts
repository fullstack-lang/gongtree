import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongtreeComponent } from './gongtree.component';

describe('GongtreeComponent', () => {
  let component: GongtreeComponent;
  let fixture: ComponentFixture<GongtreeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongtreeComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongtreeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
