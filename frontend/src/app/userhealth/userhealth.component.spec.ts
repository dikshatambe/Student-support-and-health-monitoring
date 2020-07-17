import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { UserhealthComponent } from './userhealth.component';

describe('UserhealthComponent', () => {
  let component: UserhealthComponent;
  let fixture: ComponentFixture<UserhealthComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ UserhealthComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(UserhealthComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
