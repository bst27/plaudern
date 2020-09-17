import { BrowserModule } from '@angular/platform-browser';
import {isDevMode, NgModule} from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CommentsComponent } from './components/comments/comments.component';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NavigationComponent } from './components/navigation/navigation.component';
import { LayoutModule } from '@angular/cdk/layout';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatButtonModule } from '@angular/material/button';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatIconModule } from '@angular/material/icon';
import { MatListModule } from '@angular/material/list';
import { CommentCardComponent } from './components/comment-card/comment-card.component';
import {MatCardModule} from '@angular/material/card';
import { CommentDetailsCardComponent } from './components/comment-details-card/comment-details-card.component';
import { StoreModule } from '@ngrx/store';
import {commentsReducer} from './store/comments/comments.reducer';
import { EffectsModule } from '@ngrx/effects';
import {CommentsEffects} from './store/comments/comments.effects';
import {logger} from './store/dev/logger.reducer';
import {environment} from '../environments/environment';
import { LoginComponent } from './components/login/login/login.component';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';

@NgModule({
  declarations: [
    AppComponent,
    CommentsComponent,
    NavigationComponent,
    CommentCardComponent,
    CommentDetailsCardComponent,
    LoginComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    LayoutModule,
    MatToolbarModule,
    MatButtonModule,
    MatSidenavModule,
    MatIconModule,
    MatListModule,
    MatCardModule,
    StoreModule.forRoot({
      comments: commentsReducer,
    }, {
      metaReducers: !environment.production ? [
        logger,
      ] : []
    }),
    EffectsModule.forRoot([
      CommentsEffects,
    ]),
    MatFormFieldModule,
    MatInputModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
