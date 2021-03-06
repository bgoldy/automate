import { HttpErrorResponse } from '@angular/common/http';
import { Action } from '@ngrx/store';

import { Project } from './project.model';

export enum ProjectActionTypes {
  GET_ALL         = 'PROJECT::GET_ALL',
  GET_ALL_SUCCESS = 'PROJECT::GET_ALL::SUCCESS',
  GET_ALL_FAILURE = 'PROJECT::GET_ALL::FAILURE',
  GET             = 'PROJECT::GET',
  GET_SUCCESS     = 'PROJECT::GET::SUCCESS',
  GET_FAILURE     = 'PROJECT::GET::FAILURE',
  CREATE          = 'PROJECT::CREATE',
  CREATE_SUCCESS  = 'PROJECT::CREATE::SUCCESS',
  CREATE_FAILURE  = 'PROJECT::CREATE::FAILURE',
  DELETE          = 'PROJECT::DELETE',
  DELETE_SUCCESS  = 'PROJECT::DELETE::SUCCESS',
  DELETE_FAILURE  = 'PROJECT::DELETE::FAILURE',
  UPDATE          = 'PROJECT::UPDATE',
  UPDATE_SUCCESS  = 'PROJECT::UPDATE::SUCCESS',
  UPDATE_FAILURE  = 'PROJECT::UPDATE::FAILURE'
}

export interface ProjectSuccessPayload {
  project: Project;
}

export class GetProjects implements Action {
  readonly type = ProjectActionTypes.GET_ALL;
}

export interface GetProjectsSuccessPayload {
  projects: Project[];
}

export class GetProjectsSuccess implements Action {
  readonly type = ProjectActionTypes.GET_ALL_SUCCESS;

  constructor(public payload: GetProjectsSuccessPayload) { }
}

export class GetProjectsFailure implements Action {
  readonly type = ProjectActionTypes.GET_ALL_FAILURE;

  constructor(public payload: HttpErrorResponse) { }
}

export class GetProject implements Action {
  readonly type = ProjectActionTypes.GET;

  constructor(public payload: { id: string }) { }
}

export class GetProjectSuccess implements Action {
  readonly type = ProjectActionTypes.GET_SUCCESS;

  constructor(public payload: ProjectSuccessPayload) { }
}

export class GetProjectFailure implements Action {
  readonly type = ProjectActionTypes.GET_FAILURE;

  constructor(public payload: HttpErrorResponse) { }
}

export class CreateProject implements Action {
  readonly type = ProjectActionTypes.CREATE;
  constructor(public payload: {id: string, name: string}) { }
}

export class CreateProjectSuccess implements Action {
  readonly type = ProjectActionTypes.CREATE_SUCCESS;
  constructor(public payload: ProjectSuccessPayload) { }
}

export class CreateProjectFailure implements Action {
  readonly type = ProjectActionTypes.CREATE_FAILURE;
  constructor(public payload: HttpErrorResponse) { }
}

export class DeleteProject implements Action {
  readonly type = ProjectActionTypes.DELETE;

  constructor(public payload: { id: string }) { }
}

export class DeleteProjectSuccess implements Action {
  readonly type = ProjectActionTypes.DELETE_SUCCESS;

  constructor(public payload: { id: string }) { }
}

export class DeleteProjectFailure implements Action {
  readonly type = ProjectActionTypes.DELETE_FAILURE;

  constructor(public payload: HttpErrorResponse) { }
}

export class UpdateProject implements Action {
  readonly type = ProjectActionTypes.UPDATE;

  constructor(public payload: { id: string, name: string }) { }
}

export class UpdateProjectSuccess implements Action {
  readonly type = ProjectActionTypes.UPDATE_SUCCESS;

  constructor(public payload: ProjectSuccessPayload) { }
}

export class UpdateProjectFailure implements Action {
  readonly type = ProjectActionTypes.UPDATE_FAILURE;

  constructor(public payload: HttpErrorResponse) { }
}

export type ProjectActions =
  | GetProjects
  | GetProjectsSuccess
  | GetProjectsFailure
  | GetProject
  | GetProjectSuccess
  | GetProjectFailure
  | CreateProject
  | CreateProjectSuccess
  | CreateProjectFailure
  | DeleteProject
  | DeleteProjectSuccess
  | DeleteProjectFailure
  | UpdateProject
  | UpdateProjectSuccess
  | UpdateProjectFailure;
