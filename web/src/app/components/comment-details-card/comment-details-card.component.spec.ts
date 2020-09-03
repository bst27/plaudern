import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CommentDetailsCardComponent } from './comment-details-card.component';

describe('CommentDetailsCardComponent', () => {
  let component: CommentDetailsCardComponent;
  let fixture: ComponentFixture<CommentDetailsCardComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CommentDetailsCardComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CommentDetailsCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
