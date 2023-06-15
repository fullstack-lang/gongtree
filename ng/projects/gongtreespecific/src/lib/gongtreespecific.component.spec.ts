import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongtreespecificComponent } from './gongtreespecific.component';

describe('GongtreespecificComponent', () => {
  let component: GongtreespecificComponent;
  let fixture: ComponentFixture<GongtreespecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongtreespecificComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongtreespecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
