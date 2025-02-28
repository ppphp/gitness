/*
 * Copyright 2023 Harness, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import React, { useState } from 'react'
import {
  Container,
  Layout,
  Text,
  Button,
  ButtonVariation,
  Formik,
  useToaster,
  ButtonSize,
  TextInput,
  FormInput,
  Dialog,
  StringSubstitute
} from '@harnessio/uicore'
import cx from 'classnames'
import { Color, FontVariation, Intent } from '@harnessio/design-system'
import { Icon } from '@harnessio/icons'
import { useMutate } from 'restful-react'
import { ACCESS_MODES, permissionProps, voidFn } from 'utils/Utils'
import { useStrings } from 'framework/strings'
import type { TypesRepository } from 'services/code'
import { useAppContext } from 'AppContext'
import { useGetSpaceParam } from 'hooks/useGetSpaceParam'
import { RepoVisibility } from 'utils/GitUtils'
import { useModalHook } from 'hooks/useModalHook'
import useDeleteRepoModal from './DeleteRepoModal/DeleteRepoModal'
import Private from '../../../icons/private.svg'
import css from '../RepositorySettings.module.scss'

interface GeneralSettingsProps {
  repoMetadata: TypesRepository | undefined
  refetch: () => void
}

const GeneralSettingsContent = (props: GeneralSettingsProps) => {
  const { repoMetadata, refetch } = props
  const { openModal: openDeleteRepoModal } = useDeleteRepoModal()

  const [editDesc, setEditDesc] = useState(ACCESS_MODES.VIEW)
  const { showError, showSuccess } = useToaster()

  const space = useGetSpaceParam()
  const { standalone } = useAppContext()
  const { hooks } = useAppContext()
  const { getString } = useStrings()
  const currRepoVisibility = repoMetadata?.is_public === true ? RepoVisibility.PUBLIC : RepoVisibility.PRIVATE

  const [repoVis, setRepoVis] = useState<RepoVisibility>(currRepoVisibility)
  const { mutate } = useMutate({
    verb: 'PATCH',
    path: `/api/v1/repos/${repoMetadata?.path}/+/`
  })

  const permEditResult = hooks?.usePermissionTranslate?.(
    {
      resource: {
        resourceType: 'CODE_REPOSITORY'
      },
      permissions: ['code_repo_edit']
    },
    [space]
  )
  const permDeleteResult = hooks?.usePermissionTranslate?.(
    {
      resource: {
        resourceType: 'CODE_REPOSITORY'
      },
      permissions: ['code_repo_delete']
    },
    [space]
  )
  const ModalComponent: React.FC = () => {
    return (
      <Dialog
        className={css.dialogContainer}
        style={{ width: 610, maxHeight: '95vh', overflow: 'auto' }}
        title={getString('changeRepoVis')}
        isOpen
        onClose={hideModal}>
        <Text>
          <StringSubstitute
            str={getString('changeRepoVisContent')}
            vars={{
              repoVis: <span className={css.text}>{repoVis}</span>,
              repoText:
                repoVis === RepoVisibility.PUBLIC
                  ? getString('createRepoModal.publicLabel')
                  : getString('createRepoModal.privateLabel')
            }}
          />
        </Text>
        <hr className={css.dividerContainer} />
        <Layout.Horizontal className={css.buttonContainer}>
          <Button
            margin={{ right: 'medium' }}
            type="submit"
            text={getString('confirm')}
            variation={ButtonVariation.PRIMARY}
            onClick={() => {
              mutate({ is_public: repoVis === RepoVisibility.PUBLIC ? true : false })
                .then(() => {
                  showSuccess(getString('repoUpdate'))
                  hideModal()
                })
                .catch(err => {
                  showError(err)
                })
              refetch()
            }}
          />
          <Button
            text={getString('cancel')}
            variation={ButtonVariation.TERTIARY}
            onClick={() => {
              hideModal()
            }}
          />
        </Layout.Horizontal>
      </Dialog>
    )
  }
  const [openModal, hideModal] = useModalHook(ModalComponent, [])

  return (
    <Formik
      formName="repoGeneralSettings"
      initialValues={{
        name: repoMetadata?.uid,
        desc: repoMetadata?.description,
        isPublic: currRepoVisibility
      }}
      onSubmit={voidFn(mutate)}>
      {formik => {
        return (
          <Layout.Vertical padding={{ top: 'medium' }}>
            <Container padding="large" margin={{ bottom: 'medium' }} className={css.generalContainer}>
              <Layout.Horizontal padding={{ bottom: 'medium' }}>
                <Container className={css.label}>
                  <Text color={Color.GREY_600} className={css.textSize}>
                    {getString('repositoryName')}
                  </Text>
                </Container>
                <Container className={css.content}>
                  <Text color={Color.GREY_800} className={css.textSize}>
                    {repoMetadata?.uid}
                  </Text>
                </Container>
              </Layout.Horizontal>
              <Layout.Horizontal>
                <Container className={cx(css.label, css.descText)}>
                  <Text color={Color.GREY_600} className={css.textSize}>
                    {getString('description')}
                  </Text>
                </Container>
                <Container className={css.content}>
                  {editDesc === ACCESS_MODES.EDIT ? (
                    <Layout.Horizontal className={css.editContainer}>
                      <TextInput
                        className={cx(css.textContainer, css.textSize)}
                        onChange={evt => {
                          formik.setFieldValue('desc', (evt.currentTarget as HTMLInputElement)?.value)
                        }}
                        value={formik.values.desc || repoMetadata?.description}
                        name="desc"
                      />
                      <Layout.Horizontal className={css.buttonContainer}>
                        <Button
                          className={css.saveBtn}
                          margin={{ right: 'medium' }}
                          type="submit"
                          text={getString('save')}
                          variation={ButtonVariation.SECONDARY}
                          size={ButtonSize.SMALL}
                          onClick={() => {
                            mutate({ description: formik.values?.desc })
                              .then(() => {
                                showSuccess(getString('repoUpdate'))
                              })
                              .catch(err => {
                                showError(err)
                              })
                            setEditDesc(ACCESS_MODES.VIEW)
                            refetch()
                          }}
                        />
                        <Button
                          text={getString('cancel')}
                          variation={ButtonVariation.TERTIARY}
                          size={ButtonSize.SMALL}
                          onClick={() => {
                            setEditDesc(ACCESS_MODES.VIEW)
                          }}
                        />
                      </Layout.Horizontal>
                    </Layout.Horizontal>
                  ) : (
                    <Text color={Color.GREY_800} className={css.textSize}>
                      {formik?.values?.desc || repoMetadata?.description}
                      <Button
                        className={css.textSize}
                        text={getString('edit')}
                        icon="Edit"
                        variation={ButtonVariation.LINK}
                        onClick={() => {
                          setEditDesc(ACCESS_MODES.EDIT)
                        }}
                        {...permissionProps(permEditResult, standalone)}
                      />
                    </Text>
                  )}
                </Container>
              </Layout.Horizontal>
            </Container>
            <Container padding="large" margin={{ bottom: 'medium' }} className={css.generalContainer}>
              <Layout.Horizontal padding={{ bottom: 'medium' }}>
                <Container className={css.label}>
                  <Text color={Color.GREY_600} font={{ size: 'small' }}>
                    {getString('repoVisibility')}
                  </Text>
                </Container>
                <Container className={css.content}>
                  <FormInput.RadioGroup
                    name="isPublic"
                    label=""
                    onChange={evt => {
                      setRepoVis((evt.target as HTMLInputElement).value as RepoVisibility)
                    }}
                    className={css.radioContainer}
                    items={[
                      {
                        label: (
                          <Container>
                            <Layout.Horizontal>
                              <Icon
                                className={css.iconContainer}
                                name="git-clone-step"
                                size={20}
                                margin={{ left: 'small', right: 'medium' }}
                              />
                              <Container>
                                <Layout.Vertical spacing="xsmall">
                                  <Text font={{ size: 'small' }}>{getString('public')}</Text>
                                  <Text font={{ variation: FontVariation.TINY }}>
                                    {getString('createRepoModal.publicLabel')}
                                  </Text>
                                </Layout.Vertical>
                              </Container>
                            </Layout.Horizontal>
                          </Container>
                        ),

                        value: RepoVisibility.PUBLIC
                      },
                      {
                        label: (
                          <Container>
                            <Layout.Horizontal>
                              <Container className={css.iconContainer} margin={{ left: 'small', right: 'medium' }}>
                                <img width={20} height={20} src={Private} />
                              </Container>
                              <Container margin={{ left: 'small' }}>
                                <Layout.Vertical spacing="xsmall">
                                  <Text font={{ size: 'small' }}>{getString('private')}</Text>
                                  <Text font={{ variation: FontVariation.TINY }}>
                                    {getString('createRepoModal.privateLabel')}
                                  </Text>
                                </Layout.Vertical>
                              </Container>
                            </Layout.Horizontal>
                          </Container>
                        ),
                        value: RepoVisibility.PRIVATE
                      }
                    ]}
                  />
                  <hr className={css.dividerContainer} />
                  <Layout.Horizontal className={css.buttonContainer}>
                    {repoVis !== currRepoVisibility ? (
                      <Button
                        margin={{ right: 'medium' }}
                        type="submit"
                        text={getString('save')}
                        variation={ButtonVariation.PRIMARY}
                        size={ButtonSize.SMALL}
                        onClick={() => {
                          setRepoVis(formik.values.isPublic)
                          openModal()
                        }}
                      />
                    ) : null}
                  </Layout.Horizontal>
                </Container>
              </Layout.Horizontal>
            </Container>
            <Container padding="medium" className={css.generalContainer}>
              <Container className={css.deleteContainer}>
                <Text icon="main-trash" color={Color.GREY_600} font={{ size: 'small' }}>
                  {getString('dangerDeleteRepo')}
                </Text>
                <Button
                  intent={Intent.DANGER}
                  onClick={() => {
                    openDeleteRepoModal()
                  }}
                  variation={ButtonVariation.SECONDARY}
                  text={getString('delete')}
                  {...permissionProps(permDeleteResult, standalone)}></Button>
              </Container>
            </Container>
          </Layout.Vertical>
        )
      }}
    </Formik>
  )
}

export default GeneralSettingsContent
