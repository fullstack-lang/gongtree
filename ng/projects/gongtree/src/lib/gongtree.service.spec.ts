import { TestBed } from '@angular/core/testing';

import { GongtreeService } from './gongtree.service';

describe('GongtreeService', () => {
  let service: GongtreeService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongtreeService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
